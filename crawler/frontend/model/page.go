package model

//填充到html页面的字段
type SearchResult struct {
	Hits     int64
	Start    int
	Query    string //搜索条件
	PrevFrom int
	NextFrom int
	Items    []interface{}
}
