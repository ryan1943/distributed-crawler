package persist

import (
	"context"
	"distributed-crawler/crawler/engine"
	"errors"
	"log"

	"gopkg.in/olivere/elastic.v5"
)

//把数据保存在ElasticSearch里面
//index在elasticsearch里面相当于数据库名
func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		//Must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item  #%d: %v", itemCount, item)
			itemCount++
			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

//保存数据
func Save(client *elastic.Client, index string, item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().
		Index(index). //类似数据库名
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())

	if err != nil {
		return err
	}
	//fmt.Printf("%+v", resp)
	return nil
}
