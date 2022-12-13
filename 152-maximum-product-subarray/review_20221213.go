package main

func maxProduct_20221213(nums []int) int {
	ans, maxF, minF := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		ans = max(ans, maxF)
	}

	return ans
}
