package main

import "strings"

func main() {

}

func reverseWords(s string) string {
	tmp := strings.Split(s, " ")
	res := ""
	for i := 0; i < len(tmp); i++ {
		tmp[i] = reverse(tmp[i])
		res += tmp[i] + " "
	}
	return res[:len(res)-1]
	//return strings.Trim(res, " ")
}
func reverse(s string) string {
	tmp := ""
	for i := len(s) - 1; i >= 0; i-- {
		tmp = tmp + string(s[i])
	}
	return tmp
}
