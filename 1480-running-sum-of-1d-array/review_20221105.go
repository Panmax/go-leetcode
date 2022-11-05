package main

func runningSum_20221105(nums []int) []int {
	if len(nums) <= 0 {
		return []int{}
	}
	res := make([]int, len(nums))
	res[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		res[i] = nums[i] + res[i-1]
	}
	return res
}
