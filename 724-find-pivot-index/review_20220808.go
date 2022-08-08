package main

func pivotIndex_20220808(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	// total - sum - nums[i] == sum
	sum := 0
	for i, num := range nums {
		if total == sum*2+num {
			return i
		}
		sum += num
	}
	return -1
}
