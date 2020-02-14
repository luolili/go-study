package main

func main() {

}
func findNumIn2dArray(matrix [][]int, target int) bool {
	//matrix[i,j] 作为标志位
	var i = len(matrix) - 1
	j := 0
	for i >= 0 {
		if j < len(matrix[i]) {
			if matrix[i][j] < target {
				j++
			} else if matrix[i][j] > target {
				i--
			} else {
				return true
			}
		} else {
			return false
		}

	}
	return false
}
