package main

import "sort"

func fourSum_20230505(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSumTarget_20230505(nums, 4, 0, target)
}

func nSumTarget_20230505(nums []int, n int, start int, target int) [][]int {
	if n < 2 || len(nums) < n {
		return nil
	}
	var res [][]int
	if n == 2 {
		left := start
		right := len(nums) - 1
		for left < right {
			num1 := nums[left]
			num2 := nums[right]
			if num1+num2 < target {
				for left < right && nums[left] == num1 {
					left++
				}

			} else if num1+num2 > target {
				for left < right && nums[right] == num2 {
					right--
				}
			} else {
				for left < right && nums[left] == num1 {
					left++
				}
				for left < right && nums[right] == num2 {
					right--
				}
				res = append(res, []int{num1, num2})
			}
		}
	} else {
		for i := start; i < len(nums); i++ {
			num := nums[i]
			arrs := nSumTarget_20230505(nums, n-1, i+1, target-num)
			for _, arr := range arrs {
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
