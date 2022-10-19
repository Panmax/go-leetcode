package main

func searchRange_20221019(nums []int, target int) []int {
	left := search_20221019(nums, target, true)
	right := search_20221019(nums, target, false) - 1
	// left <= right
	if left <= right && right < len(nums) && nums[left] == target && nums[right] == target {
		return []int{left, right}
	}
	return []int{-1, -1}
}

func search_20221019(nums []int, target int, low bool) int {
	left, right := 0, len(nums)-1
	res := len(nums)
	// left <= right
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] > target || (low && nums[mid] >= target) {
			right = mid - 1
			res = mid
		} else {
			left = mid + 1
		}
	}
	return res
}
