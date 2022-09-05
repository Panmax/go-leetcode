package main

func subarraySum_20220905(nums []int, k int) int {
	res := 0
	m := make(map[int]int)
	m[0] = 1
	preSum := 0
	for _, num := range nums {
		preSum += num
		res += m[preSum-k]
		m[preSum]++
	}
	return res
}
