package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"
	res := isUnique(str)
	fmt.Println(res)
}

func isUnique(astr string) bool {
	//遍历 string 获得里面的字符
	for _, ch1 := range astr {
		s := string(ch1)
		index1 := strings.Index(astr, s)
		index2 := strings.LastIndex(astr, s)
		if index1 != index2 {
			return false
		}
	}
	return true
}
func firstUniqueChar(astr string) byte {
	//遍历 string 获得里面的字符
	for _, ch1 := range astr {
		s := string(ch1)
		index1 := strings.Index(astr, s)
		index2 := strings.LastIndex(astr, s)
		if index1 == index2 {
			return byte(ch1)
		}
	}
	return ' '
}
