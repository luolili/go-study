package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	strReader := strings.NewReader("hello world")

	bufReader := bufio.NewReader(strReader)

	data, _ := bufReader.Peek(7)
	//bufReader.Buffered() 缓存了多少个字符，这里把hello world缓存了 11
	fmt.Println(data, bufReader.Buffered())
	str, _ := bufReader.ReadString(' ')
	fmt.Println(str, bufReader.Buffered()) //5

	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "hu")     //写入到F (file)
	fmt.Fprint(w, "fgs ft") //写入到F (file)
	w.Flush()               //输出到stdout

}
