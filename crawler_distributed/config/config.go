package config

const (
	//Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	//Service port
	ItemSaverPort = 1234
	WorkerPort0   = 9000

	//ES的index名，类似数据库名
	ElasticIndex = "dating_profile"

	//要注册和本地调用的rpc服务
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	//Rate limiting
	Qps = 20
)
