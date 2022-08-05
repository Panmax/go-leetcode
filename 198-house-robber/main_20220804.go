package main

// dp[i] = max(dp[i-1], dp[i-2]+nums[i])
func rob_20220804(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dps := make([]int, len(nums))
	dps[0] = nums[0]
	dps[1] = max(dps[0], dps[1])
	for i := 2; i < len(nums); i++ {
		dps[i] = max(dps[i-1], dps[i-2]+nums[i])
	}
	return dps[len(nums)-1]
}
