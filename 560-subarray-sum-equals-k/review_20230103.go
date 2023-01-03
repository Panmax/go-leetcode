package main

func subarraySum_20230103(nums []int, k int) int {
	var res, sum int
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
