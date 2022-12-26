package main

func maximum_20221226(a int, b int) int {
	// uint64
	index := uint64(a-b) >> 63
	return []int{a, b}[index]
}
