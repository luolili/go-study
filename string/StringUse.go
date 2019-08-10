package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello wo"
	fmt.Println(strings.Contains(str, "wo"))
	fmt.Println(strings.Index(str, "l")) //2
	fmt.Println(strings.HasPrefix(str, "hel"))
	fmt.Println(strings.HasSuffix(str, "lr"))

	ss := "1#2#21"
	splitStr := strings.Split(ss, "#")

	//fmt.Println(splitStr)
	fmt.Println(strings.Join(splitStr, "#"))
	//----转换
	fmt.Println("---转换---")
	fmt.Println(strconv.Itoa(12))
	fmt.Println(strconv.Atoi("321"))
	//strconv.Atoi: parsing "3s21": invalid syntax
	//fmt.Println(strconv.Atoi("3s21"))

	//把字符串转为bool
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("T"))
	fmt.Println(strconv.ParseFloat("3.5621332", 32)) //3.562133312225342 <nil>
	fmt.Println(strconv.ParseFloat("3.5621332", 64)) //3.5621332 <nil>
	//把bool转为字符串
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(21, 2)) //10101
	fmt.Println(strconv.FormatInt(21, 10))
	//fmt.Println(strconv.FormatFloat(21.21,32))

}
