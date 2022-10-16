package main

import "sort"

func fourSum_20221016(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSumTarget_20221016(nums, 4, 0, target)
}

func nSumTarget_20221016(nums []int, n int, start int, target int) [][]int {
	var res [][]int
	if n < 2 || len(nums) < n {
		return res
	}
	// n == 2, 不是 len(nums)==2
	if n == 2 {
		left, right := start, len(nums)-1
		for left < right {
			num1, num2 := nums[left], nums[right]
			sum := num1 + num2
			if sum == target {
				// 加进去的是 num1 和 num2
				res = append(res, []int{num1, num2})
				for left < right && nums[left] == num1 {
					left++
				}
				for left < right && nums[right] == num2 {
					right--
				}
			} else if sum < target {
				for left < right && nums[left] == num1 {
					left++
				}
			} else {
				for left < right && nums[right] == num2 {
					right--
				}
			}
		}
	} else {
		for i := start; i < len(nums); i++ {
			num := nums[i]
			tmpArrs := nSumTarget_20221016(nums, n-1, i+1, target-num)
			for _, arr := range tmpArrs {
				arr = append(arr, num)
				res = append(res, arr)
			}
			for i+1 < len(nums) && nums[i+1] == num {
				i++
			}
		}
	}
	return res
}
