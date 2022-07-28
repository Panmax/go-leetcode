package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	dict1 := make(map[byte]byte) // s -> t
	dict2 := make(map[byte]byte) // t -> s
	for i := range s {
		x := s[i]
		y := t[i]
		if dict1[x] > 0 || dict2[y] > 0 {
			if dict1[x] != y || dict2[y] != x {
				return false
			}
			continue
		}
		dict1[x] = y
		dict2[y] = x
	}
	return true
}

// https://leetcode.cn/problems/isomorphic-strings/
func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("badc", "baba"))
}
