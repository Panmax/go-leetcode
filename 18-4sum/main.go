package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSumTarget(nums, 4, 0, target)
}

func nSumTarget(nums []int, n int, start int, target int) [][]int {
	var res [][]int
	if n < 2 || len(nums) == 0 || len(nums) < n {
		return res
	}
	if n == 2 {
		left, right := start, len(nums)-1
		for left < right {
			num1, num2 := nums[left], nums[right]
			sum := num1 + num2
			if sum < target {
				for left < right && nums[left] == num1 {
					left++
				}
			} else if sum > target {
				for left < right && nums[right] == num2 {
					right--
				}
			} else {
				res = append(res, []int{nums[left], nums[right]})
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
			nRes := nSumTarget(nums, n-1, i+1, target-num)
			for _, arr := range nRes {
				arr = append(arr, num)
				res = append(res, arr)
			}
			// i+1 和 num 比较，外边有一次++了，这里不用多加
			for i < len(nums) && nums[i+1] == num {
				i++
			}
		}
	}
	return res
}
