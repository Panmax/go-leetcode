package main

func longestConsecutive_20230409(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}
	var longest int
	for _, num := range nums {
		length := 1
		// !m[num-1]
		if !m[num-1] {
			curNum := num
			for {
				curNum += 1
				if m[curNum] {
					length++
				} else {
					break
				}
			}
		}
		if length > longest {
			longest = length
		}
	}
	return longest
}
