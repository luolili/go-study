package main

func main() {

}
func findRepeatNum(nums []int) int {
	var len = len(nums)
	for i := 0; i < len; i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			var temp = nums[i]
			nums[i] = nums[temp]
			nums[temp] = temp
		}

	}
	return -1
}
