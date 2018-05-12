这是单机版的并发爬虫，数据存储在docker 里面的elasticsearch
使用命令如下:
docker run -d-p 9200:9200 elasticsearch
映射到物理目录
 docker run -p 9200:9200 -v c:/data/elastic:/usr/share/elasticsearch/data -d elasticsearch

访问方式:
localhost:9200/dating_profile/zhenai/_search?size=100&q=男 已购房 已购车 Age:(<30) Height([178 TO 185])

