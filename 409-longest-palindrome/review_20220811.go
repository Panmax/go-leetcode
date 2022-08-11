package main

func longestPalindrome_20220811(s string) int {
	length := 0
	slot := [128]int{}
	for _, c := range s {
		slot[c]++
	}

	for _, count := range slot {
		length += count / 2 * 2
		if (count%2 == 1) && (length%2 == 0) {
			length++
		}
	}

	return length
}
