package main

func subarraySum_20220919(nums []int, k int) int {
	res := 0
	m := make(map[int]int)
	sum := 0
	m[0] = 1
	for _, num := range nums {
		sum += num
		res += m[sum-k]
		m[sum]++
	}
	return res
}
