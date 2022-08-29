package main

func runningSum_20220829(nums []int) []int {
	sums := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			sums[0] = nums[0]
		} else {
			sums[i] = sums[i-1] + nums[i]
		}
	}

	return sums
}
