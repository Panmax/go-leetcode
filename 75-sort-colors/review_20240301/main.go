package main

func sortColors(nums []int) {
	n := sortNum(nums, 0)
	sortNum(nums[n:], 1)
}

func sortNum(nums []int, num int) int {
	var count int
	for i, n := range nums {
		if n == num {
			nums[i], nums[count] = nums[count], nums[i]
			count++
		}
	}
	return count
}
