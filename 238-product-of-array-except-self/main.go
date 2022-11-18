package main

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	// 前缀积
	for i := 1; i < len(nums); i++ {
		// nums[i-1]
		res[i] = res[i-1] * nums[i-1]
	}
	// 后缀积
	tmp := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= tmp
		tmp *= nums[i]
	}

	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	println(productExceptSelf(nums))
}
