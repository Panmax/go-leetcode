package main

func subsets_20231213(nums []int) [][]int {
	n := len(nums)
	var res [][]int
	for mask := 0; mask < 1<<n; mask++ {
		var set []int
		for i, num := range nums {
			if mask>>i&1 == 1 {
				set = append(set, num)
			}
		}

		res = append(res, set)
	}
	return res
}
