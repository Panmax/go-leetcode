package main

func numSubarraysWithSum_20230510(nums []int, goal int) int {
	m := make(map[int]int)
	// 别忘了这个
	m[0] = 1
	var ans, sum int
	for _, num := range nums {
		sum += num
		ans += m[sum-goal]
		m[sum]++
	}
	return ans
}
