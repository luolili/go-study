package main

func main() {

}
func maxValueV2(grid [][]int) int {
	dp := make([][]int, len(grid)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[0])+1)
	}
	dp[0][0] = grid[0][0]

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			dp[i+1][j+1] = grid[i][j] + getMax(dp[i][j+1], dp[i+1][j])

		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
func maxValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {

			} else if i == 0 {
				grid[i][j] = grid[i][j] + grid[i][j-1]
			} else if j == 0 {
				grid[i][j] = grid[i][j] + grid[i-1][j]
			} else {
				grid[i][j] = grid[i][j] + getMax(grid[i][j-1], grid[i-1][j])
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
