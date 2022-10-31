package main

import "sort"

func fourSum_20221031(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSumTarget_20221031(nums, 4, 0, target)
}

func nSumTarget_20221031(nums []int, n int, start int, target int) [][]int {
	res := make([][]int, 0)
	if n < 2 || len(nums) < n {
		return res
	}
	if n == 2 {
		left := start
		right := len(nums) - 1
		for left < right {
			num1, num2 := nums[left], nums[right]
			if num1+num2 < target {
				for left < right && nums[left] == num1 {
					left++
				}
			} else if num1+num2 > target {
				for left < right && nums[right] == num2 {
					right--
				}
			} else {
				res = append(res, []int{num1, num2})
				for left < right && nums[left] == num1 {
					left++
				}
				for left < right && nums[right] == num2 {
					right--
				}
			}
		}

	} else {
		for i := start; i < len(nums); i++ {
			num := nums[i]
			arrs := nSumTarget_20221031(nums, n-1, i+1, target-num)
			for _, arr := range arrs {
				arr = append(arr, num)
				res = append(res, arr)
			}
			// for
			for i+1 < len(nums) && nums[i+1] == num {
				i++
			}
		}
	}
	return res
}
