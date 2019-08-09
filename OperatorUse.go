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
}
