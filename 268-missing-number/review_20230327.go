package main

import "sort"

func missingNumber_20230327(nums []int) int {
	sort.Ints(nums)
	for i, num := range nums {
		if i != num {
			return i
		}
	}
	return len(nums)
}
