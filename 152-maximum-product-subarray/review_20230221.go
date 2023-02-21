package main

func maxProduct_20230221(nums []int) int {
	ans, maxVal, minVal := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxVal, minVal
		maxVal = max(nums[i]*mx, max(nums[i], nums[i]*mn))
		minVal = min(nums[i]*mx, min(nums[i], nums[i]*mn))
		ans = max(ans, maxVal)
	}
	return ans
}
