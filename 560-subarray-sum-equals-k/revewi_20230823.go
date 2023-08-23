package main

func subarraySum_20230823(nums []int, k int) int {
	var res int
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
