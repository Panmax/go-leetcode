package main

func rob_20230417(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dps := make([]int, n)
	dps[0] = nums[0]
	dps[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dps[i] = max(dps[i-1], dps[i-2]+nums[i])
	}
	return dps[n-1]
}
