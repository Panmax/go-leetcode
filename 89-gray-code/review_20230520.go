package main

// 看不懂的话去看这个题解：
// https://leetcode.cn/problems/gray-code/solution/gray-code-jing-xiang-fan-she-fa-by-jyd/
func grayCode_2023050(n int) []int {
	var res []int
	res = append(res, 0)
	base := 1
	for i := 0; i < n; i++ {
		size := len(res)
		for j := size - 1; j >= 0; j-- {
			res = append(res, base+res[j])
		}
		base = base << 1
	}

	return res
}
