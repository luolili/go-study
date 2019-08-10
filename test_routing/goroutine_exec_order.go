package main

import (
	"fmt"
	"time"
)

func PrintExecTime() {
	fmt.Println(time.Now())

}
func main() {
	for i := 0; i < 20; i++ {
		go PrintExecTime() //并行执行，打印出来的时间基本一样
	}
	time.Sleep(1 * time.Second)
	fmt.Println("i am main routine here")

	//sync.Mutex{} //加锁
}
