package main //必须是main包才可以运行main方法

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {

	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}

	return p, err
}

func sampleReadFromString() {

	strings.NewReader("from string")
	data, _ := ReadFrom(strings.NewReader("from string"), 12)
	fmt.Println(data)

}

func sampleReadFromCtrl() {

	fmt.Println("please input something")
	data, _ := ReadFrom(os.Stdin, 11)
	fmt.Println(string(data))

}

func sampleReadFromFile() {
	file, _ := os.Open("main.go")
	defer file.Close()
	data, _ := ReadFrom(file, 222)
	fmt.Println(string(data)) //把byte 转为string
}
func main() {
	//sampleReadFromString()
	//sampleReadFromCtrl()
	sampleReadFromFile()
}
