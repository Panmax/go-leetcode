package main

func singleNumber_20231112(nums []int) int {
	var n int
	for _, num := range nums {
		n ^= num
	}
	return n
}
