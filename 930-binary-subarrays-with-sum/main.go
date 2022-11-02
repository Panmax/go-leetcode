package main

func numSubarraysWithSum(nums []int, goal int) int {
	sum := 0
	m := make(map[int]int)
	ans := 0
	m[0] = 1
	for _, num := range nums {
		sum += num
		ans += m[sum-goal]
		m[sum]++
	}

	return ans
}
