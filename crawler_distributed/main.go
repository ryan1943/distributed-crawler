package main

import (
	"distributed-crawler/crawler/engine"
	"distributed-crawler/crawler/scheduler"
	"distributed-crawler/crawler/zhenai/parser"
	"distributed-crawler/crawler_distributed/config"
	itemsaver "distributed-crawler/crawler_distributed/persist/client"
	"distributed-crawler/crawler_distributed/rpcsupport"
	worker "distributed-crawler/crawler_distributed/worker/client"
	"flag"
	"log"
	"net/rpc"
	"strings"

	"github.com/garyburd/redigo/redis"
)

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
	redisConn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer redisConn.Close()

	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor := worker.CreateProcessor(pool) //worker的客户端

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      50,
		ItemChan:         itemChan,
		RequestProcessor: processor,
		RedisConn:        redisConn,
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
