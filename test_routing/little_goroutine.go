package main

import (
	"fmt"
	"time"
)

func little_goroutine() {
	fmt.Println("little go routine func")

}
func main() {
	//创建一个协程
	go little_goroutine()
	time.Sleep(1 * time.Second)
	fmt.Println("i am a main go routine")
}
