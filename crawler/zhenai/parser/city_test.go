package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	contents, err := ioutil.ReadFile("city_test.html")
	if err != nil {
		panic(err)
	}
	result := ParseCity(contents, "")

	expectedUrls := []string{
		"http://album.zhenai.com/u/108415017",
		"http://album.zhenai.com/u/1314495053",
		"http://album.zhenai.com/u/110171680",
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url %d: %s; but was %s",
				i, url, result.Requests[i].Url)
		}
	}

}
