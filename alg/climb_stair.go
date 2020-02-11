package main

import "fmt"

func main() {
	stair := climbStair(5)

	fmt.Println(stair)
}

func climbStair(n int) int {
	if n == 1 {
		return 1
	}
	var dp = make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
