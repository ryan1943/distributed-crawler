package main

import (
	"fmt"

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

	username, err := redis.String(c.Do("GET", "key"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
}

func Save(c redis.Conn, key string) string {
	_, err := c.Do("GET", key)
	if err == nil {
		return "0"
	} else {
		_, err = c.Do("SET", key, "1")
		if err != nil {
			fmt.Println("redis set failed:", err)
		}
	}
}
