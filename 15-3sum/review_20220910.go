package main

import "sort"

func threeSum_20220910(nums []int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0)
	for i := range nums {
		if nums[i] > 0 {
			break
		}
		// 需要加这个判断
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
