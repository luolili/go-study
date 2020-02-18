package main

func main() {

}
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= (i >> 1); j++ {
			dp[i] = getMax(dp[i], getMax(j, dp[j])*getMax(i-j, dp[i-j]))
		}
	}
	return dp[n]
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
