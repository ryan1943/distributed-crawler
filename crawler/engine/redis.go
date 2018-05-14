package engine

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

//去掉重复的解析过的url
func isDuplicate(c redis.Conn, key string, value string) bool {
	v, err := c.Do("GET", key)
	if v == nil || err != nil {
		_, err = c.Do("SET", key, value)
		if err != nil {
			log.Println("redis set failed:", err)
		}
		return false
	} else {
		return true //已经存在
	}

}
