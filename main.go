package main

import (
	"fmt"
	"show"
)

func main() {
	fmt.Println("halo", 2)
	str := fmt.Sprintf("float %f", 2.1235)
	fmt.Println(str)
	fmt.Printf("%v\n", Data{}) //dd
	show.Show()                //调用自定义的包
}

type Data struct {
}

func (self Data) String() string {
	return "dd"

}
