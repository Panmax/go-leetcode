package main

func grayCode(n int) []int {
	res := make([]int, 0, n*n)
	res = append(res, 0)
	head := 1
	for i := 0; i < n; i++ {
		size := len(res)
		for j := size - 1; j >= 0; j-- {
			res = append(res, head+res[j])
		}
		head = head << 1
	}
	return res
}
