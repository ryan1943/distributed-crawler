数据存储在docker的elasticsearch里面

ES启动时两种方法：

第一种临时存储数据:

docker run -d -p 9200:9200 elasticsearch

第二种是映射到物理目录可以保存在本地

docker run -d -p 9200:9200 -v /home/chenwd/data/elastic:/usr/share/elasticsearch/data elasticsearch

启动redis并将主机中当前data目录挂载到容器的/data：

docker run -d -p 6379:6379 -v /home/chenwd/data/redis:/data  redis redis-server --appendonly yes

分布式爬虫需要启动rpc服务：

启用保存数据到ES的rpc服务:
go run crawler_distributed/persist/server/itemsaver.go --port=1234

启用两个爬取工作的rpc服务:

go run crawler_distributed/worker/server/worker.go --port=9000

go run crawler_distributed/worker/server/worker.go --port=9001

启动主程序：go run crawler_distributed/main.go --itemsaver_host=":1234" --worker_hosts=":9000,:9001"

工作过程中若30s内没有接收到解析结果则超时停止

前端展示需要启动http服务：

go run crawler/frontend/starter.go
本地访问：localhost 进入搜索首页

临时网站：123.207.88.187


![001](https://github.com/ryan1943/distributed-crawler/blob/master/images/001.png)

搜索结果

![002](https://github.com/ryan1943/distributed-crawler/blob/master/images/002.png)




