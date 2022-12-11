package main

func searchRange_20221211(nums []int, target int) []int {
	left := search_20221211(nums, target, true)
	right := search_20221211(nums, target, false) - 1
	if left <= right && right < len(nums) && nums[left] == target && nums[right] == target {
		return []int{left, right}
	}
	return []int{-1, -1}
}

func search_20221211(nums []int, target int, low bool) int {
	res := len(nums)
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target || (low && nums[mid] >= target) {
			right = mid - 1
			res = mid
		} else {
			left = mid + 1
		}
	}

	return res
}
