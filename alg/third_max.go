package main

import "sort"

func main() {

}

func thirdMax(nums []int) int {
	var len = len(nums)
	if len == 1 {
		return nums[0]
	}
	if len == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}
	sort.Ints(nums)
	if nums[len-1] != nums[len-2] {
		for i := len - 3; i >= 0; i-- {

			if nums[0] == nums[len-2] {
				return nums[len-1]
			}
			if nums[i] != nums[len-2] {
				return nums[i]
			}

		}
	}
	if nums[len-1] == nums[len-2] {
		var n = 1
		for i := len - 3; i >= 0; i-- {
			if nums[i] != nums[len-2] {
				n = i
				break
			}
		}
		for i := n; i >= 0; i-- {
			if nums[0] == nums[n] {
				return nums[len-1]
			}
			if nums[i] != nums[n] {
				return nums[i]
			}
		}
	}

	return nums[len-3]
}
