package main

func findLucky_20230115(arr []int) int {
	ans := -1
	m := make(map[int]int)
	for _, val := range arr {
		m[val]++
	}
	for k, v := range m {
		if k == v && v > ans {
			ans = v
		}
	}
	return ans
}
