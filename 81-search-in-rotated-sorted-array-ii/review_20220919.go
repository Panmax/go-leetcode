package main

func search_20220919(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return nums[0] == target
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return true
		} else if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[left] <= nums[mid] { // 小于等于
			if target < nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}
