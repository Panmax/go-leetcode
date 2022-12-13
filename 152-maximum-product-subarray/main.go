package main

func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF

		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))

		ans = max(ans, maxF)
	}
	return ans
}

func maxProduct_BigSpace(nums []int) int {
	length := len(nums)
	maxF := make([]int, length)
	minF := make([]int, length)
	for i, num := range nums {
		maxF[i] = num
		minF[i] = num
	}
	for i := 1; i < length; i++ {
		maxF[i] = max(maxF[i-1]*nums[i], max(nums[i], minF[i-1]*nums[i]))
		minF[i] = min(minF[i-1]*nums[i], min(nums[i], maxF[i-1]*nums[i]))
	}
	res := nums[0]
	for _, val := range maxF {
		res = max(res, val)
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
