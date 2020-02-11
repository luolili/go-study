package main

import "math"

func main() {

}
func drinkLast(n int) string {
	var names = []string{"A", "B", "C", "D"}
	var base = len(names)
	// 表示有几个自己
	var i = 0
	for n > base {
		n -= base
		// copy itself
		base *= 2
		i++
	}
	var index = int(math.Ceil(float64(n / int(math.Pow(2, float64(i))))))
	return names[index-1]
}
