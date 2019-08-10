package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
) //引入go redis

const (
	REDIS_KEY = "mobile:p30:n3000:1"
)

func main() {
	//链接redis
	conn, err := redis.Dial("tcp", "106.12.205.88:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close() //关闭redis链接：最后执行
	//执行redis命令:左边的变量不要和之前的一样
	_, err01 := conn.Do("lpush", REDIS_KEY, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	if err01 != nil {
		fmt.Println(err01)
		return
	}

}
