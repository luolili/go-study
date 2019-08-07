//import "fmt"
package main //默认的包：package关键字必须是在go文件最开始的位置
//导入的包
import "fmt"

//定义常量
const NAME = "GO"

//加类型的声明方法
const CAR string = "ff"

//全局var
var userId = "dd g"
var dog string = "dd g"

//自定义的函数
func sayHalo() {
	fmt.Print("ha")

}

/*
多行注释
默认的main方法
*/
func main() {
	fmt.Println("halo")
	fmt.Print(userId)
	fmt.Print(NAME)
	sayHalo() //函数的调用

}
