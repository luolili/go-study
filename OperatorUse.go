package main

import "fmt"

func main() {
	a := 2
	a++
	fmt.Println(a) //3
	b := 5
	b--
	fmt.Println(b)     //4
	fmt.Println(a - b) //-1
	fmt.Println(a + b) //7
	fmt.Println(a * b) //12
	fmt.Println(a / b) //0
	fmt.Println(a % b) //3
	//--关系运算符
	fmt.Println("--关系运算符")
	fmt.Println(a == b) //false
	fmt.Println(a >= b) //false
	fmt.Println(a <= b) //true
	fmt.Println(a != b) //true
	fmt.Println(a > b)  //false
	fmt.Println(a < b)  //true
	fmt.Println("--logic运算符")
	fmt.Println(true && false) //false
	fmt.Println(true && true)  //true
	fmt.Println(true || false) //tru
	fmt.Println(!true)
	fmt.Println("--按位运算符")
	//& 都为1-->1 ; | 有一个1-->1 ；^ 不同-->1; >>  <<

	n := 3
	//n>>1
	fmt.Println(n >> 1) //1
	fmt.Println(n << 1) //6
	x1 := byte(0)
	x2 := byte(0)
	fmt.Println(x1 & x2) //0
	fmt.Println(x1 | x2) //0
	fmt.Println(x1 ^ x2) //0

}
