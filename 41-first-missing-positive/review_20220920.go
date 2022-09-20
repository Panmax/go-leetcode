package main

func firstMissingPositive_20220920(nums []int) int {
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
		if nums[i] < 1 || nums[i] > n {
			// 这里是 1
			nums[i] = 1
		}
	}
	for i := range nums {
		// 改变的是 a 的元素
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
