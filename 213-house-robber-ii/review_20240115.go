package main

func rob_20240115(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(_rob_20240115(nums[:n-1]), _rob_20240115(nums[1:]))
}

func _rob_20240115(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1]) // max!!!
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}
