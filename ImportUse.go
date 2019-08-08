package main

import (
	format "fmt" //fmt的别名format
	. "show"     //在依赖包前面加. 可以在调用这个包里面的方法的时候不写包名
	_ "show02"   // _ 只会执行init方法
)

func main() {
	format.Println("halo") //取了别名之后就只能用别名了
	//show.Show()
	//show.Show()
	Show()
	//show02.Show02()//
}
