package main

func subsets_20230916(nums []int) [][]int {
	n := len(nums)
	var ans [][]int

	for mask := 0; mask < 1<<n; mask++ {
		var set []int
		for i, num := range nums {
			if mask>>i&1 == 1 {
				set = append(set, num)
			}
		}
		ans = append(ans, set)
	}
	return ans
}
