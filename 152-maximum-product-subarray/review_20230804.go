package main

func maxProduct_20230804(nums []int) int {
	ans, maxVal, minVal := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxVal, minVal
		maxVal = max(nums[i], max(nums[i]*mx, nums[i]*mn))
		minVal = min(nums[i], min(nums[i]*mx, nums[i]*mn))
		ans = max(ans, maxVal)
	}
	return ans
}
