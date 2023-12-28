package main

func longestConsecutive_20231228(nums []int) int {
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		m[num] = true
	}
	var longest int
	for _, num := range nums {
		var length int
		if !m[num-1] {
			length = 1
			currNum := num
			// currNum+1
			for m[currNum+1] {
				length++
				currNum++
			}
			if length > longest {
				longest = length
			}
		}
	}
	return longest
}
