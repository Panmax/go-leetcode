package main

import "strings"

func reverseWords_20230518(s string) string {
	ss := strings.Split(s, " ")
	var slice []string
	for _, str := range ss {
		if str != "" {
			slice = append(slice, str)
		}
	}
	for i := 0; i < len(slice)/2; i++ {
		slice[i], slice[len(slice)-1-i] = slice[len(slice)-1-i], slice[i]
	}
	return strings.Join(slice, " ")
}
