package main

func firstMissingPositive_20230706(nums []int) int {
	find1 := false
	for _, num := range nums {
		if num == 1 {
			find1 = true
			break
		}
	}
	if !find1 {
		return 1
	}
	for i := range nums {
		if nums[i] <= 0 || nums[i] > len(nums) {
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
	return len(nums) + 1
}
