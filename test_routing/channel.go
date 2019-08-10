package main

import "fmt"

//quit哎channel里面标记当前协程已经完成
func Sum(var1, var2 int, quit chan bool) {
	z := var1 + var2
	fmt.Println(z)
	quit <- true //向channel里面写quit为true

}
func main() {
	channels := make([]chan bool, 10)

	for i := 0; i < 10; i++ {
		channels[i] = make(chan bool)
		go Sum(i, i, channels[i])
	}

}
