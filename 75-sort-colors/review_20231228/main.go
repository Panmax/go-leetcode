package main

func sortColors(nums []int) {
	n := swipeNum(nums, 0)
	swipeNum(nums[n:], 1)
}

func swipeNum(nums []int, num int) int {
	var count int
	for i, n := range nums {
		if n == num {
			nums[i], nums[count] = nums[count], nums[i]
			count++
		}
	}
	return count
}
