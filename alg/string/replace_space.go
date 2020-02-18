package main

import (
	"fmt"
	"strings"
)

// 格式化：%b 二进制；%c:unicode;%d 十进制；%x:16进制
//float32:8位指数，23位尾数
func main() {
	str := "helo wo"
	replace := strings.Replace(str, " ", "u", 2)
	fmt.Println(replace)

}
func replaceSpace(s string) string {
	for _, ch := range s {
		if ' ' == ch {
			strings.Replace(s, " ", "%20", strings.Count(s, string(' ')))
		}
	}
	return s
}
