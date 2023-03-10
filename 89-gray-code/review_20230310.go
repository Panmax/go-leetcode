package main

func grayCode_20230310(n int) []int {
	var res []int
	res = append(res, 0)
	head := 1
	for i := 0; i < n; i++ {
		size := len(res)
		for j := size - 1; j >= 0; j-- {
			res = append(res, head+res[j])
		}
		head <<= 1
	}
	return res
}
