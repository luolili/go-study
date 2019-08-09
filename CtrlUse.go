package main

import (
	"fmt"
	"time"
)

//控制语句
func main() {
	a := 2
	if a < 0 {
		fmt.Println("res: a<0")
	} else { //注意：else 必须紧跟着}, 不可换行
		fmt.Println("res: a>0")
	}

	//---switch: 输出CCC
	switch 3 {
	case 1:
		fmt.Println("---")

	case 2:
		fmt.Println("+++")
	default:
		fmt.Println("CCC")

	}

	//输出int
	var b interface{}
	b = 22
	switch b.(type) {
	case int:
		fmt.Print("int")
	case string:
		fmt.Print("string")
	default:
		fmt.Print("b")

	}
	//for
	/*for {
		fmt.Println("oi")
		time.Sleep(2*time.Second)
	}
	*/
	//循环9次
	for i := 1; i < 10; i++ {
		fmt.Println("oi")
	}

	m := []string{"app", "ban", "org"}
	for key, value := range m {
		fmt.Println(key)   //从0开始的索引
		fmt.Println(value) //元素值
	}

	for _, value := range m {
		//fmt.Println(key)//从0开始的索引
		fmt.Println(value) //元素值
	}

	goto One
	fmt.Println("---") //被跳过
One:
	fmt.Println("cc")

	//break
	for {
		fmt.Println("p")
		time.Sleep(2 * time.Second)
		break //只打印了一个p
	}

	for i := 1; i < 3; i++ {
		for j := 1; j < 3; j++ {
			fmt.Println("vv")
			time.Sleep(2 * time.Second)
			break
		}
	}

	for i := 1; i < 3; i++ {
		if i >= 1 {
			fmt.Println("ff")
			time.Sleep(1 * time.Second)
			continue //不会执行fmt.Println("last")
		}
		fmt.Println("last")
	}

}
