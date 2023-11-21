package main

func grayCode_20231121(n int) []int {
	var res []int
	res = append(res, 0)
	base := 1
	for i := 0; i < n; i++ {
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, res[j]+base)
		}
		base <<= 1
	}

	return res
}
