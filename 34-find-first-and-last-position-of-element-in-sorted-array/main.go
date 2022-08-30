package main

func searchRange(nums []int, target int) []int {
	leftIndex := search(nums, target, true)
	rightIndex := search(nums, target, false) - 1
	if leftIndex <= rightIndex && rightIndex < len(nums) && nums[leftIndex] == target && nums[rightIndex] == target {
		return []int{leftIndex, rightIndex}
	}
	return []int{-1, -1}
}

func search(nums []int, target int, low bool) int {
	left, right := 0, len(nums)-1
	// 这里是 nums 的长度
	ans := len(nums)
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] > target || (low && nums[mid] >= target) {
			right = mid - 1
			ans = mid
		} else {
			left = mid + 1
		}
	}

	return ans
}
