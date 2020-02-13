package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("d")
		// 恢复 panic
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("e")
	}()
	f()
}
func f() {

	fmt.Println("a")
	// stop 后面的代码不执行
	panic(33)
	fmt.Println("b")
	fmt.Println("c")

}
