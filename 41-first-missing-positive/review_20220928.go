package main

func firstMissingPositive_20220928(nums []int) int {
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
		// || nums[i] > len(nums)
		if nums[i] <= 0 || nums[i] > len(nums) {
			nums[i] = 1
		}
	}
	for i := range nums {
		// -1
		a := Abs(nums[i]) - 1
		nums[a] = -Abs(nums[a])
	}
	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}
