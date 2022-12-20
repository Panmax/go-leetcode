package main

func maximum_20221220(a int, b int) int {
	index := uint64(a-b) >> 63
	return []int{a, b}[index]
}
