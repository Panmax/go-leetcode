package main

func longestPalindrome_20220815(s string) int {
	charMap := make(map[int32]int)
	for _, c := range s {
		charMap[c]++
	}
	length := 0
	for _, count := range charMap {
		length += count / 2 * 2
		if count%2 == 1 && length%2 == 0 {
			length++
		}
	}
	return length
}
