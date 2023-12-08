package main

func singleNumber_20231208(nums []int) int {
	var n int
	for _, num := range nums {
		n ^= num
	}
	return n
}
