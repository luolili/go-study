package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := make([]int, 3)
	arr[0] = 2
	arr[1] = 3
	arr[2] = 1
	sort.Ints(arr)
	for _, i := range arr {
		fmt.Println(i)
	}

}
func getLeastNum(arr []int, k int) []int {
	res := make([]int, k)
	// asc
	sort.Ints(arr)
	for i := 0; i < k; i++ {
		res[i] = arr[i]
	}
	return res
}
