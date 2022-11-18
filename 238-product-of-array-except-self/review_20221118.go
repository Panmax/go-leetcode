package main

func productExceptSelf_20221118(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}
	tmp := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= tmp
		tmp *= nums[i]
	}
	return res
}
