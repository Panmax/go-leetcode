package main

func pivotIndex_20230509(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	sum := 0
	for i, num := range nums {
		if sum*2+num == total {
			return i
		}
		sum += num
	}
	return -1
}
