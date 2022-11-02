package main

func numSubarraysWithSum_20221102(nums []int, goal int) int {
	m := make(map[int]int)
	ant := 0
	m[0] = 1
	sum := 0
	for _, num := range nums {
		sum += num
		ant += m[sum-goal]
		m[sum]++
	}
	return ant
}
