package main

func nextPermutation_20220906(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	// 这里要判断 i>=0
	if i >= 0 {
		j := len(nums) - 1
		for j > i && nums[i] >= nums[j] {
			j--
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	// i+1 后的区间反转
	reverse_20220906(nums[i+1:])
}

func reverse_20220906(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}
