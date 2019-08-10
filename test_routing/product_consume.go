package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//链接redis
	conn, err := redis.Dial("tcp", "106.12.205.88:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer conn.Close()//关闭redis链接：最后执行
	item, err02 := redis.String(conn.Do("lpop", REDIS_KEY))
	if err02 != nil {
		fmt.Println(err02)
		return
	}

	fmt.Printf("res: %s", item) //res:10
}
