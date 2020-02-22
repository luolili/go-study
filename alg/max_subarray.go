package main

func main() {

}

//入一个整型数组，数组里有正数也有负数。
// 数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值
func maxSubarray(nums []int) int {
	res := nums[0]
	sum := 0
	for _, v := range nums {
		sum = getMax(sum+v, v)
		res = getMax(res, sum)
	}
	return res
}
