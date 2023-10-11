package main

func findLucky_20231011(arr []int) int {
	ans := -1
	m := make(map[int]int, len(arr))
	for _, v := range arr {
		m[v]++
	}
	for k, v := range m {
		// && v > ans
		if k == v && v > ans {
			ans = v
		}
	}
	return ans
}
