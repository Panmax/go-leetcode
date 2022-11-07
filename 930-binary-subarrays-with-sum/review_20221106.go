package main

func numSubarraysWithSum_20221106(nums []int, goal int) int {
	ant := 0
	m := make(map[int]int)
	sum := 0
	m[0] = 1
	for _, num := range nums {
		sum += num
		ant += m[sum-goal]
		m[sum]++
	}
	return ant
}
