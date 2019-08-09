package main

import "fmt"

/*
iota是用于常量的定义
*/
const a = iota
const (
	b    = iota
	_           //让c=2
	_           //让c=3
	c    = iota //自增1
	d, e = iota, iota + 3
	f, g
	h = iota
)

func main() {
	fmt.Println("a: ")
	fmt.Println(a)      //0
	fmt.Println(b)      //0
	fmt.Println(c)      // 1
	fmt.Println("----") // 1
	fmt.Println(d)      // 4
	fmt.Println(e)      // 7
	fmt.Println(f)      // 5
	fmt.Println(g)      // 8
	fmt.Println(h)      // 6

}
