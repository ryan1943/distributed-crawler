package view

import (
	"distributed-crawler/crawler/frontend/model"
	"html/template"
	"io"
)

type SearchResultView struct {
	template *template.Template
}

//工厂函数
func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{
		template: template.Must(template.ParseFiles(filename)),
	}
}

//data数据填充到模板
//把渲染后的模板写入到writer
func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}
