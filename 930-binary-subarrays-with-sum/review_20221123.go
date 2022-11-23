package main

func numSubarraysWithSum_20221123(nums []int, goal int) int {
	m := make(map[int]int)
	m[0] = 1
	ans := 0
	sum := 0
	for _, num := range nums {
		sum += num
		ans += m[sum-goal]
		m[sum]++
	}
	return ans
}
