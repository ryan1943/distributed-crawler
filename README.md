学习分布式爬虫的项目代码，数据存储在docker的elasticsearch里面
docker的命令:启动时docker run -d -p 9200:9200 elasticsearch 查看进程docker ps
映射到物理目录
 docker run -p 9200:9200 -v c:/data/elastic:/usr/share/elasticsearch/data -d elasticsearch
