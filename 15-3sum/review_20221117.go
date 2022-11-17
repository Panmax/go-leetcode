package main

import "sort"

func threeSum_20221117(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)
	for i, num := range nums {
		if num > 0 {
			break
		}
		if i > 0 && nums[i-1] == num {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			if num+nums[left]+nums[right] == 0 {
				res = append(res, []int{num, nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if num+nums[left]+nums[right] < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return res
}
