package main

func subarraySum_20221013(nums []int, k int) int {
	res, sum := 0, 0
	m := make(map[int]int)
	m[0] = 1
	for _, num := range nums {
		sum += num
		if count, ok := m[sum-k]; ok {
			res += count
		}
		m[sum]++
	}
	return res
}
