package main

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return true
		}

		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[left] <= nums[mid] { // 左边有序
			if nums[left] <= target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右边有序
			if nums[right] >= target && nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}
