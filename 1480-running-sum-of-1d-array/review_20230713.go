package main

func runningSum_20230713(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		res[i] = nums[i] + res[i-1]
	}

	return res
}
