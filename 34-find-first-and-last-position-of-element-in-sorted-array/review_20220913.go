package main

func searchRange_20220913(nums []int, target int) []int {
	indexLow := search_20220913(nums, target, true)
	indexHigh := search_20220913(nums, target, false) - 1
	// 还需要判断 indexHigh < len(nums)
	if indexLow <= indexHigh && indexHigh < len(nums) && nums[indexLow] == target && nums[indexHigh] == target {
		return []int{indexLow, indexHigh}
	}
	return []int{-1, -1}
}

func search_20220913(nums []int, target int, low bool) int {
	left, right := 0, len(nums)-1
	// len(nums)
	res := len(nums)
	for left <= right {
		mid := (right + left) / 2
		// 前边是大于，后边要加上 或 (low && nums[mid] >= target)
		if nums[mid] > target || (low && nums[mid] >= target) {
			right = mid - 1
			// 别忘了下边
			res = mid
		} else {
			left = mid + 1
		}
	}
	return res
}
