package main

func maximum_20231122(a int, b int) int {
	index := uint64(a-b) >> 63
	return []int{a, b}[index]
}
