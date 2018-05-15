package engine

import (
	"time"

	"log"

	"github.com/garyburd/redigo/redis"
)

//启动引擎
type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
	RedisConn        redis.Conn
}

type Processor func(Request) (ParseResult, error)

//调度器
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request //生成worker chan
	Run()
}

type ReadyNotifier interface {
	//发送worker chan 到scheduler的workerChan，它的类型是 chan chan engine.Request
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()
	//生成相应数目的worker chan，并发送到scheduler的worker队列去
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r) //发送request到scheduler的request队列去
	}

	for {
		//30s超时退出
		tm := time.After(30 * time.Second)
		var result ParseResult
		select {
		case <-tm:
			log.Println("30s timeout.No task execution!")
			return
		case result = <-out:
		}
		//result := <-out
		for _, item := range result.Items {
			go func() {
				e.ItemChan <- item
			}()
		}
		for _, request := range result.Requests {
			if isDuplicate(e.RedisConn, request.Url, "1") {
				log.Println("duplicate url")
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func (e *ConcurrentEngine) createWorker(
	in chan Request,
	out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			//发送到worker队列去,完成一次工作后再次放回到worker队列去
			ready.WorkerReady(in)
			//接收到这里传来的数据 activeWorker <- activeRequest, 可能要等待
			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

/*var visitedUrls = make(map[string]bool)

//去重
func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}*/
