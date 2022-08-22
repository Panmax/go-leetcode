package main

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)

	res := make([][]int, 0)
	for i, num := range nums {
		if num > 0 {
			return res
		}
		if i > 0 && nums[i-1] == num {
			continue
		}

		left := i + 1
		right := n - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				res = append(res, []int{num, nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left += 1
				}
				for left < right && nums[right] == nums[right-1] {
					right -= 1
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
