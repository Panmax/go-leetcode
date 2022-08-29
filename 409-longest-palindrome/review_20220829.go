package main

func longestPalindrome_20220829(s string) int {
	charMap := make(map[byte]int)
	for i := range s {
		charMap[s[i]] += 1
	}
	length := 0
	for _, count := range charMap {
		length += count / 2 * 2
		if length%2 == 0 && count%2 == 1 {
			length += 1
		}
	}
	return length
}
