package main

func maxProduct_20221220(nums []int) int {
	ans, maxV, minV := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxV, minV
		maxV = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minV = min(mn*nums[i], min(nums[i], mx*nums[i]))
		ans = max(ans, maxV)
	}
	return ans
}
