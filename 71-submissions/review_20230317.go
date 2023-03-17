package main

import "strings"

func simplifyPath_20230317(path string) string {
	var stack []string
	paths := strings.Split(path, "/")
	for _, p := range paths {
		if p == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if p != "" && p != "." {
			stack = append(stack, p)
		}
	}
	return "/" + strings.Join(stack, "/")
}
