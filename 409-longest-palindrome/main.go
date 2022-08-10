package main

func longestPalindrome(s string) int {
	slot := [128]int{}
	for i := 0; i < len(s); i++ {
		slot[s[i]]++
	}

	length := 0
	for _, count := range slot {
		length += count / 2 * 2
		if count%2 == 1 && length%2 == 0 {
			length++
		}
	}
	return length
}
