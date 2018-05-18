package main

import (
	"learncrawler/crawler/frontend/controller"
	"net/http"
)

func main() {
	//加上这行view目录下的css和js文件才能访问
	//开启了一个资源访问的文件服务器
	http.Handle("/", http.FileServer(
		http.Dir("crawler/frontend/view")))
	http.Handle("/search", controller.CreateSearchResultHandler(
		"crawler/frontend/view/template1.html"))
	err := http.ListenAndServe(":8888", nil) //使用DefaultServeMux作为服务器的主handler
	if err != nil {
		panic(err)
	}
}
