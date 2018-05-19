package main

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "key", "1")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	//测试去重功能
	username, err := redis.String(c.Do("GET", "key"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	res := Save(c, "key6", "1")
	if res {
		fmt.Println("没重复")
	} else {
		fmt.Println("重复了")
	}
}

func Save(c redis.Conn, key string, value string) bool {
	v, err := c.Do("GET", key)
	if v == nil || err != nil {
		_, err = c.Do("SET", key, value)
		if err != nil {
			log.Println("redis set failed:", err)
		}
		return true
	} else {
		return false //已经存在
	}

}
