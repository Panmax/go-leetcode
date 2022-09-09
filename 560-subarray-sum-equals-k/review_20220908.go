package main

func subarraySum_20220908(nums []int, k int) int {
	ans := 0
	m := make(map[int]int)
	m[0] = 1
	sum := 0
	for _, num := range nums {
		sum += num
		if m[sum-k] > 0 {
			ans += m[sum-k]
		}
		m[sum]++
	}
	return ans
}
