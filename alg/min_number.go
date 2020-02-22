package main

import "strconv"

func main() {

}
func minNum(nums []int) string {
	res := ""
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if bigger(nums[i], nums[j]) == nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	for i := 0; i < len(nums); i++ {
		res += strconv.Itoa(nums[i])
	}
	return res
}

func bigger(a, b int) int {
	// 数字转string
	aStr, bStr := strconv.Itoa(a), strconv.Itoa(b)
	abRes := aStr + bStr
	baRes := bStr + aStr
	if abRes > baRes {
		return a
	} else {
		return b
	}
}
