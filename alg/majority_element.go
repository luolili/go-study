package main

func main() {

}

// 摩尔投票法：出现次数最多的元素
func majorityElementV2(nums []int) int {
	target := 0
	count := 0
	for _, v := range nums {
		if count == 0 {
			target = v
			count++
		} else if v == target {
			count++
		} else {
			count--
		}
	}
	return target
}
func majorityElement(nums []int) int {
	target := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if target == nums[i] {
			count++
		} else {
			count--
		}
		if count == 0 {
			target = nums[i]
			count = 1
		}
	}
	return target
}
