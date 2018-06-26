package parser

import (
	"distributed-crawler/crawler/engine"
	"distributed-crawler/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test.html")
	if err != nil {
		panic(err)
	}
	result := parseProfile(contents, "http://album.zhenai.com/u/1314495053", "风中的蒲公英")

	if len(result.Items) != 1 {
		t.Errorf("Items should be contain 1 element; but was %v", len(result.Items))
	}

	actual := result.Items[0]

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1314495053",
		Type: "zhenai",
		Id:   "1314495053",
		Payload: model.Profile{
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

	if actual != expected {
		t.Errorf("expected user %v; but was %v", expected, actual)
	}
}
