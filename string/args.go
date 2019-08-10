package main

import (
	"flag"
	"fmt"
)

func main() {
	//fmt.Println(os.Args)
	methodPtr := flag.String("method", "default", "method of sample")

	valuePtr := flag.Int("value", 2, "method of sample")
	flag.Parse()
	fmt.Println(*methodPtr, *valuePtr) //default 2

}
