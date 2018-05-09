package parser

import (
	"learncrawler/crawler/engine"
	"regexp"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)

//根据城市列表页面解析各个城市的url
func ParseCityList(contents []byte, _ string) engine.ParseResult {
	matches := cityListRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	//m[1]是城市对应的url，m[2]是城市名
	for _, m := range matches {
		//result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,  //后续的解析器
		})
	}

	return result
}
