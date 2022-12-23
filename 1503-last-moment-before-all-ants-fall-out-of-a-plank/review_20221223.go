package main

func getLastMoment_20221223(n int, left []int, right []int) int {
	var ans int
	for _, v := range left {
		ans = max(ans, v)
	}
	for _, v := range right {
		ans = max(ans, n-v)
	}
	return ans
}
