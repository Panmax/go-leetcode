package main

func firstMissingPositive_20220917(nums []int) int {
	n := len(nums)
	find1 := false
	for i := range nums {
		if nums[i] == 1 {
			find1 = true
			break
		}
	}
	if !find1 {
		return 1
	}
	for i := range nums {
		if nums[i] <= 1 || nums[i] > n {
			nums[i] = 1
		}
	}
	for i := range nums {
		index := Abs(nums[i]) - 1
		nums[index] = -Abs(nums[index])
	}
	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return n + 1
}
