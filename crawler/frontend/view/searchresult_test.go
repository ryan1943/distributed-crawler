package view

import (
	"learncrawler/crawler/engine"
	"learncrawler/crawler/frontend/model"
	common "learncrawler/crawler/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template1.html")
	outFile, err := os.Create("template.test.html")
	data := model.SearchResult{}
	data.Hits = 123
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1314495053",
		Type: "zhenai",
		Id:   "1314495053",
		Payload: common.Profile{
			Name:       "风中的蒲公英",
			Gender:     "女",
			Age:        41,
			Height:     158,
			Weight:     48,
			Income:     "3001-5000元",
			Marriage:   "离异",
			Education:  "中专",
			Occupation: "公务员",
			Hokou:      "四川阿坝",
			Xinzuo:     "处女座",
			House:      "已购房",
			Car:        "未购车",
		},
	}
	for i := 0; i < 10; i++ {
		data.Items = append(data.Items, item)
	}
	err = view.Render(outFile, data)
	if err != nil {
		panic(err)
	}

}
