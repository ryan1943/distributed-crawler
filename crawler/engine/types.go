package engine

//解析器函数体
type ParserFunc func(contents []byte, url string) ParseResult

//请求
type Request struct {
	Url        string
	ParserFunc ParserFunc
}

//解析器返回的结果
type ParseResult struct {
	Requests []Request
	Items    []Item
}

//存储的一条记录
type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
