package main

import "fmt"

func isSubsequence(s, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

func main() {
	fmt.Println(isSubsequence("abc", "uabc"))
}
