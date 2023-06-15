package main

func longestConsecutive_20230615(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}
	var longest int
	for _, num := range nums {
		var longth int
		if !m[num-1] {
			longth = 1
			currNum := num
			for m[currNum+1] {
				currNum++
				longth++
			}
			if longth > longest {
				longest = longth
			}
		}
	}

	return longest
}
