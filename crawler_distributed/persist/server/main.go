package main

import (
	"fmt"
	"learncrawler/crawler_distributed/persist"
	"learncrawler/crawler_distributed/rpcsupport"
	"log"

	"learncrawler/crawler_distributed/config"

	"gopkg.in/olivere/elastic.v5"
)

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort),
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
