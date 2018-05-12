package main

import (
	"flag"
	"learncrawler/crawler/engine"
	"learncrawler/crawler/scheduler"
	"learncrawler/crawler/zhenai/parser"
	"learncrawler/crawler_distributed/config"
	itemsaver "learncrawler/crawler_distributed/persist/client"
	"learncrawler/crawler_distributed/rpcsupport"
	worker "learncrawler/crawler_distributed/worker/client"
	"log"
	"net/rpc"
	"strings"
)

//启用itemsaver的rpc端口: go run crawler_distributed/persist/server/itemsaver.go --port=1234
//启用worker的rpc端口: go run crawler_distributed/worker/server/worker.go --port=9000
//					   go run crawler_distributed/worker/server/worker.go --port=9001
//go run crawler_distributed/main.go --itemsaver_host=":1234" --worker_hosts=":9000,:9001"
var (
	itemSaverHost = flag.String(
		"itemsaver_host", "", "itemsaver host")
	workerHosts = flag.String(
		"worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()

	itemChan, err := itemsaver.ItemSaver(*itemSaverHost) //客户端
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor := worker.CreateProcessor(pool) //worker的客户端

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      50,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

}

//爬虫连接池
//根据启用的rpc端口数去生成相应的rpc客户端数
//假如hosts的数量是2，engine的WorkerCount是50
//本机50个worker的goroutine去轮流调用这2个rpc客户端
//每个rpc客户端对应的服务端各自使用了25个goroutine去并发处理
func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
