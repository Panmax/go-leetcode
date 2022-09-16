package main

func firstMissingPositive_20220916(nums []int) int {
	n := len(nums)
	find1 := false
	for i := range nums {
		if nums[i] == 1 {
			find1 = true
		}
	}
	if !find1 {
		return 1
	}

	for i := range nums {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = 1
		}
	}

	for i := range nums {
		a := Abs(nums[i]) - 1
		nums[a] = -Abs(nums[a])
	}

	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}
