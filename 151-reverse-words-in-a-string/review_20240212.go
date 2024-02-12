package main

import "strings"

func reverseWords_20240211(s string) string {
	ss := strings.Split(s, " ")
	var slice []string
	for i := range ss {
		if ss[i] != "" {
			slice = append(slice, ss[i])
		}
	}
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
	return strings.Join(slice, " ")
}
