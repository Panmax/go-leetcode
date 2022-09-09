package main

// 132, 213
func nextPermutation_20220908(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	// 这里判断 i >= 0
	if i >= 0 {
		j := len(nums) - 1
		// 这里判断 nums[i] >= nums[j]
		for j > i && nums[i] >= nums[j] {
			j--
		}
		if j > i {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	reverse_20220908(nums[i+1:])
}

func reverse_20220908(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
