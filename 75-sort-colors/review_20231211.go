package main

func sortColors_20231211(nums []int) {
	count := swipeColors_20231211(nums, 0)
	swipeColors_20231211(nums[count:], 1)
}

func swipeColors_20231211(nums []int, target int) int {
	var count int
	for i, num := range nums {
		if num == target {
			nums[count], nums[i] = nums[i], nums[count]
			count++
		}
	}
	return count
}
