package main

func nextPermutation_20220913(nums []int) {
	n := len(nums)
	i := n - 2
	// i >=0 && nums[i] >= nums[i+1]
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		// nums[i] >= nums[j]
		for j > i && nums[i] >= nums[j] {
			j--
		}
		if j > i {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	reverse_20220913(nums[i+1:])
}

func reverse_20220913(nums []int) {
	// n = len(nums)
	for i, n := 0, len(nums); i < n/2; i++ {
		// nums[n-1-i]
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
