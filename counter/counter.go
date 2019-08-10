package main

import (
	"bufio"
	"fmt"
	"os"
)

//计算文件行数
func main() {
	if len(os.Args) < 2 {
		return
	}

	var filename = os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close() //在程序结束之后关闭file
	reader := bufio.NewReader(file)
	var line int
	for {
		_, isPrefix, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
			break
		}

		if !isPrefix {
			line++
		}
	}

	fmt.Println(line)
}
