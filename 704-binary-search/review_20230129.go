package main

func search_20230129(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			// -1
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
