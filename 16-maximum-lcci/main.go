package main

func maximum(a int, b int) int {
	arr := []int{a, b}
	i := uint64(a-b) >> 63
	return arr[i]
}

func main() {
	println(maximum(2, 3))
}
