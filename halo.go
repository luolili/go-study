package main //默认的包

import "fmt"

//定义常量
const NAME = "GO"

//全局var
var userId = "dd g"

/*
多行注释
默认的main方法
*/
func main() {
	fmt.Println("halo")
	fmt.Print(userId)
	fmt.Print(NAME)

}
