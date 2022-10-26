package main

func pivotIndex_20221026(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	sum := 0
	// total - num = sum*2
	for i, num := range nums {
		if total-num == sum*2 {
			return i
		}
		sum += num
	}

	return -1
}
