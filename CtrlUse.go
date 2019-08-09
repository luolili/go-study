package main

import (
	"fmt"
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
}
