package main

import (
	"fmt"
	"unsafe"
)

/**
不可以定义没有用到的var
有这么多占不同内存的数据类型是为了高效利用内存
类型的零值： 是一个var被声明之后的默认值
*/
func main() {
	//-- int
	var i uint8 = 2
	var i02 uint16 = 2
	var i03 uint32 = 2
	var i04 uint = 2
	fmt.Println(unsafe.Sizeof(i))   // 占用1个字节=8bit
	fmt.Println(unsafe.Sizeof(i02)) // 占用2个字节= 2 * 8bit
	fmt.Println(unsafe.Sizeof(i03)) // 占用4个字节= 4 * 8bit
	//根据操作系统而定
	fmt.Println(unsafe.Sizeof(i04)) // 占用8个字节= 8 * 8bit

	var i05 rune = 2
	fmt.Println(unsafe.Sizeof(i05)) //4
	//-- bool
	var a bool
	fmt.Println(a) //false

	var t int
	var t1 float32
	var t2 string
	fmt.Println("int 默认值：")
	fmt.Println(t) //0

	fmt.Println("float32 默认值：")
	fmt.Println(t1) //0

	fmt.Println("string 默认值：")
	fmt.Println(t2 == "") //true
	//给类型int32取别名xc
	type xc int32
	var ty xc = 23
	fmt.Println(ty)

}
