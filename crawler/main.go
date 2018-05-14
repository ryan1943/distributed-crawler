package main

import (
	"learncrawler/crawler/engine"
	"learncrawler/crawler/persist"
	"learncrawler/crawler/scheduler"
	"learncrawler/crawler/zhenai/parser"

	"github.com/garyburd/redigo/redis"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	redisConn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer redisConn.Close()
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      50,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
		RedisConn:        redisConn,
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
	// e.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	// 	ParserFunc: parser.ParseCity,
	// })
}
