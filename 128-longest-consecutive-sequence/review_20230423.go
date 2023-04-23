package main

func longestConsecutive_20230423(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}
	var longest int
	for _, num := range nums {
		if !m[num-1] {
			length := 1
			curNum := num
			for {
				if m[curNum+1] {
					length++
					curNum++
				} else {
					break
				}
			}
			if length > longest {
				longest = length
			}
		}
	}

	return longest
}
