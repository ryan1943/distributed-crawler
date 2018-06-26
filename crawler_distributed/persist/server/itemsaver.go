package main

import (
	"distributed-crawler/crawler_distributed/persist"
	"distributed-crawler/crawler_distributed/rpcsupport"
	"flag"
	"fmt"
	"log"

	"distributed-crawler/crawler_distributed/config"

	"gopkg.in/olivere/elastic.v5"
)

var port = flag.Int("port", 0, "the port for me to listen on")

//ItemSaver服务端注册
func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	/*log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort),
	config.ElasticIndex))*/
	//通过命令行来启用itemsaver的rpc端口
	//go run itemsaver.go --help帮助
	//go run itemsaver.go --port=1234
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port),
		config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
