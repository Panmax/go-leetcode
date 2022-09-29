package main

import "sort"

func threeSum_20220930(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)
	sort.Ints(nums)
	for i := range nums {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		left := i + 1
		right := n - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// left 和 left+1比，right 和 right-1比，不是和 i 比
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return res
}
