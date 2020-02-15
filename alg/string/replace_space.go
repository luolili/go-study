package main

import (
	"fmt"
	"strings"
)

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
