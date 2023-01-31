package main

func findLucky_20230130(arr []int) int {
	ans := -1
	m := make(map[int]int, 0)
	for _, val := range arr {
		m[val]++
	}
	for k, v := range m {
		if k == v && k > ans {
			ans = k
		}
	}
	return ans
}
