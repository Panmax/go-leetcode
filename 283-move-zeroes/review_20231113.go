package main

func moveZeroes_20231113(nums []int) {
	var j int
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
