package main

import (
	"fmt"
	"reflect"
)

//只支持字符串，数字，bool
const NAME = "ki"

const HU string = "ff bg"

const PI float32 = 3.14

const app, ban, num = "fr", "te", 21
const applen = len(app)
const IS_OPEN = true

func main() {

	var (
		i int32
		j string = "ko"
	)
	fmt.Println(i)
	fmt.Println(j)
	//全局变量必须用var 关键字； 首字母小写是本包的私有var
	var a, b = 2, 3
	fmt.Println(a + b)

	var c = 1.2
	fmt.Println(reflect.TypeOf(c)) //float64

	//省略var 加: 只能用在函数里面
	dd, ff := 3, 5
	fmt.Println(dd)
	fmt.Println(ff)
	//go只有显示的类型转换
	var age = 11
	age01 := float32(age)
	fmt.Println(age01)
	fmt.Println(reflect.TypeOf(age01)) //float32
	fmt.Println(applen)                //2
}
