package main

import "fmt"

func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	// total - sum - nums[i] == sum
	// sum*2 + nums[i] == total
	sum := 0
	for i := 0; i < len(nums); i++ {
		if sum*2+nums[i] == total {
			return i
		}
		sum += nums[i]
	}
	return -1
}

// https://leetcode.cn/problems/find-pivot-index/
func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums) == 3)

	nums = []int{1, 2, 3}
	fmt.Println(pivotIndex(nums) == -1)

	nums = []int{2, 1, -1}
	fmt.Println(pivotIndex(nums) == 0)
}
