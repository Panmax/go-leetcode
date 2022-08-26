package main

import "sort"

func threeSum_20220826(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i, num := range nums {
		if num > 0 {
			return res
		}
		// 判断重复
		if i > 0 && nums[i-1] == num {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := num + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{num, nums[left], nums[right]})
				// 判断重复
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				// 改变 left 和 right
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return res
}
