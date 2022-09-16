package main

func firstMissingPositive(nums []int) int {
	n := len(nums)
	find1 := false
	for i := range nums {
		if nums[i] == 1 {
			find1 = true
		}
	}
	if !find1 {
		return 1
	}

	for i := range nums {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = 1 // 等于 1
		}
	}
	for i := range nums {
		a := Abs(nums[i]) - 1
		nums[a] = -Abs(nums[a])
	}
	for i := range nums {
		if nums[i] > 0 { // 大于0
			return i + 1 // 需要+1
		}
	}
	return n + 1
}

func Abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func main() {
	print(firstMissingPositive([]int{1, 2, 0}))
}
