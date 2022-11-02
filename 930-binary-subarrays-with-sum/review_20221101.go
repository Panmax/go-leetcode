package main

func numSubarraysWithSum_20221101(nums []int, goal int) int {
	m := make(map[int]int)
	sum := 0
	ans := 0
	m[0] = 1
	for _, num := range nums {
		sum += num
		ans += m[sum-goal]
		m[sum]++
	}
	return ans
}
