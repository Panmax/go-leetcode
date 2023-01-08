package main

func numSubarraysWithSum_20230108(nums []int, goal int) int {
	m := make(map[int]int)
	m[0] = 1
	var ans, sum int
	for _, num := range nums {
		sum += num
		// sum-goal
		ans += m[sum-goal]
		m[sum]++
	}
	return ans
}
