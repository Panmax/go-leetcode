package main

import "strings"

func simplifyPath_20221202(path string) string {
	var stack []string
	slices := strings.Split(path, "/")
	for _, c := range slices {
		if c == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if c != "" && c != "." { // else if
			stack = append(stack, c)
		}
	}
	return "/" + strings.Join(stack, "/")
}
