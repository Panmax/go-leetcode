package main

func singleNumber(nums []int) int {
	var num int
	for _, n := range nums {
		num ^= n
	}
	return num
}
