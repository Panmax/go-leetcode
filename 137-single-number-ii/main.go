package main

import "sort"

func singleNumber(nums []int) int {
	var ans int32
	for i := 0; i < 32; i++ {
		var total int32
		for _, num := range nums {
			total += (int32(num) >> i) & 1
		}
		if total%3 == 1 {
			ans |= 1 << i
		}
	}
	return int(ans)
}

func singleNumber_me(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 3 {
		if i > len(nums)-2 || nums[i] != nums[i+2] {
			return nums[i]
		}
	}
	return 0
}
