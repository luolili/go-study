package main

import "fmt"

func main() {

	var a int8
	var b int8
	// 类型必须相同，才可比较，没有隐士的 类型转换
	fmt.Println(a == b)
	// 复合类型：数组，结构体
	//引用类型：地址；切片/map 不可比较
	//接口类型：动态类型+动态值
}
