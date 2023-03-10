package main

import "sort"

func missingNumber_20230310(nums []int) int {
	sort.Ints(nums)
	for i, num := range nums {
		if i != num {
			return i
		}
	}
	return len(nums)
}
