package main

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}
	var longest int
	for _, num := range nums {
		var length int
		if !m[num-1] {
			currentNum := num
			// length = 1
			length = 1
			for m[currentNum+1] {
				currentNum++
				length++
			}
		}
		if length > longest {
			longest = length
		}
	}

	return longest
}
