package main

import "fmt"

//输入一个整数，输出该数二进制表示中 1 的个数。例如，
// 把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。
func main() {
	hammingWeightV2(5)
}
func hammingWeight(n uint32) int {
	var count uint32 = 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return int(count)
}
func hammingWeightV2(num uint32) int {
	count := 0
	numStr := fmt.Sprintf("%b", num)
	fmt.Println("str: " + numStr)
	for _, ch := range numStr {
		if ch == '1' {
			count++
		}
	}
	return count
}
