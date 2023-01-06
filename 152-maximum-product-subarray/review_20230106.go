package main

func maxProduct_20230106(nums []int) int {
	ans, maxV, minV := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxV, minV
		maxV = max(nums[i]*mx, max(nums[i], nums[i]*mn))
		minV = min(nums[i]*mn, min(nums[i], nums[i]*mx))
		ans = max(ans, maxV)
	}
	return ans
}
