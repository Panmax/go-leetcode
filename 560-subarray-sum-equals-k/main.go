package main

func subarraySum(nums []int, k int) int {
	preSum := 0
	res := 0
	m := make(map[int]int)
	m[0] = 1
	for i := range nums {
		preSum += nums[i]
		if _, ok := m[preSum-k]; ok {
			res += m[preSum-k]
		}
		m[preSum]++
	}
	return res
}
