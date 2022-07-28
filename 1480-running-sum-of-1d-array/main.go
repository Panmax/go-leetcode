package main

import "fmt"

func runningSum(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	results := make([]int, len(nums))
	results[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		results[i] = results[i-1] + nums[i]
	}
	return results
}

// https://leetcode.cn/problems/running-sum-of-1d-array/
func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(runningSum(nums))

	nums = []int{1, 1, 1, 1, 1}
	fmt.Println(runningSum(nums))

	nums = []int{3, 1, 2, 10, 1}
	fmt.Println(runningSum(nums))
}
