package main

func longestPalindrome_20221009(s string) int {
	charMap := make(map[byte]int)
	for i := range s {
		charMap[s[i]]++
	}
	var res int
	for _, v := range charMap {
		res += v / 2 * 2
		if v%2 == 1 && res%2 == 0 {
			res += 1
		}
	}
	return res
}
