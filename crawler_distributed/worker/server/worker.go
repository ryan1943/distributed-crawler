package main

import (
	"flag"
	"fmt"
	"learncrawler/crawler_distributed/rpcsupport"
	"learncrawler/crawler_distributed/worker"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a prot")
		return
	}
	//log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort0),
	//	worker.CrawlService{}))
	//通过命令行来启用worker的rpc端口
	//go run worker.go --help帮助
	//go run worker.go --port=9000
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}
