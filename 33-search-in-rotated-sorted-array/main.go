package main

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	left, right := 0, n-1
	for left <= right {
		// 注意：这里是 +
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		}

		// 左半边有序，注意：小于等于
		if nums[0] <= nums[mid] {
			// 注意：两边和 target 做等于查询
			if nums[mid] > target && nums[0] <= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右半边有序
			if nums[mid] < target && nums[n-1] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
