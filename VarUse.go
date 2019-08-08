package main

import "fmt"

func main() {

	var (
		i int32
		j string
	)
	fmt.Println(i)
	fmt.Println(j)
	//全局变量必须用var 关键字； 首字母小写是本包的私有var
	var a, b = 2, 3
	fmt.Println(a + b)
	//go只有显示的类型转换
}
