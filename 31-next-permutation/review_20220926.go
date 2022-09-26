package main

func nextPermutation_20220926(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j > i && nums[i] >= nums[j] {
			j--
		}
		if j > i {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	reverse(nums[i+1:])
}
