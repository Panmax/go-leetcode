package main

func searchRange_20220902(nums []int, target int) []int {
	left := searchBinary_20220902(nums, target, true)
	// 需要 -1 ！！
	right := searchBinary_20220902(nums, target, false) - 1
	if left <= right && right < len(nums) && nums[left] == target && nums[right] == target {
		return []int{left, right}
	}
	return []int{-1, -1}
}

func searchBinary_20220902(nums []int, target int, low bool) int {
	left := 0
	right := len(nums) - 1
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
