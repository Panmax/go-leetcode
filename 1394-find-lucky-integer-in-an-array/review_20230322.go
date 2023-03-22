package main

func findLucky_20230322(arr []int) int {
	ans := -1
	m := make(map[int]int)
	for _, num := range arr {
		m[num]++
	}
	for k, v := range m {
		if k == v && k > ans {
			ans = k
		}
	}
	return ans
}
