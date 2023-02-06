package main

func nextPermutation_20230206(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j > i && nums[i] >= nums[j] {
			j--
		}
		if j > i {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	reverse(nums[i+1:])
}
