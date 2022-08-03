package main

// dps[i] = max(dps[i-1], dps[i-2]+nums[i])
func rob_20220803(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	dps := make([]int, length)
	dps[0] = nums[0]
	dps[1] = max(nums[0], nums[1])
	for i := 2; i < length; i++ {
		dps[i] = max(dps[i-1], dps[i-2]+nums[i])
	}
	return dps[length-1]
}
