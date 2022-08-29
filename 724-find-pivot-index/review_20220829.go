package main

func pivotIndex_20220829(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	// total - sum - nums[i] = sum
	// total == 2*sum+nums[i]
	sum := 0
	for i, num := range nums {
		if total == 2*sum+num {
			return i
		}
		sum += num
	}
	return -1
}
