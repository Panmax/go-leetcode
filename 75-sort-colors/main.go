package main

func sortColors(nums []int) {
	// 先把0放前边
	count := swipeColors(nums, 0)
	// 再把1放前边
	swipeColors(nums[count:], 1)
}

func swipeColors(nums []int, target int) int {
	var count int
	for i, num := range nums {
		if num == target {
			nums[count], nums[i] = nums[i], nums[count]
			count++
		}
	}
	return count
}
