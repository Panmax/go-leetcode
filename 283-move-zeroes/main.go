package main

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	j := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
