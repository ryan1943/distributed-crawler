数据存储在docker的elasticsearch里面
ES启动时两种方法：
第一种临时存储: docker run -d -p 9200:9200 elasticsearch
第二种是映射到物理目录可以保存
docker run -p 9200:9200 -v c:/data/elastic:/usr/share/elasticsearch/data -d elasticsearch
redis启动：
docker run -p 6379:6379 -v c:/data/redis:/data  -d redis:3.2 redis-server --appendonly yes

分布式爬虫需要启动rpc服务：
启用itemsaver的rpc服务: go run crawler_distributed/persist/server/itemsaver.go --port=1234
启用两个worker的rpc服务: go run crawler_distributed/worker/server/worker.go --port=9000
					       go run crawler_distributed/worker/server/worker.go --port=9001

主程序：go run crawler_distributed/main.go --itemsaver_host=":1234" --worker_hosts=":9000,:9001"