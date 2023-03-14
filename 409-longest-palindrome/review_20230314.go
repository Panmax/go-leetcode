package main

func longestPalindrome_20230314(s string) int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	res := 0
	for _, v := range m {
		res += v / 2 * 2
		if v%2 == 1 && res%2 == 0 {
			res += 1
		}
	}
	return res
}
