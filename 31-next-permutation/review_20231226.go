package main

func nextPermutation_20231226(nums []int) {
	i := len(nums) - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		j := len(nums) - 1
		for i < j {
			if nums[j] > nums[i] {
				break
			}
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]

	}
	reverse(nums[i+1:])
}
