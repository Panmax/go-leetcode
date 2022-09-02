package main

func search_20220902(nums []int, target int) int {
	// 注意边界判断
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		}

		// 左半边有序
		// 应该是小于等于！
		if nums[0] <= nums[mid] {
			if nums[0] <= target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && nums[len(nums)-1] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
