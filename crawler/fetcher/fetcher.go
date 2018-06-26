package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"distributed-crawler/crawler_distributed/config"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

//限制访问网站的速度
var rateLimiter = time.Tick(time.Second / config.Qps)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	log.Printf("Fetching url %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong status code: %d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

}

//确认编码
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024) //Peek 返回缓存的一个切片，该切片引用缓存中前 n 个字节的数据
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
