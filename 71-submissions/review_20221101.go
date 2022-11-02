package main

import "strings"

func simplifyPath_20221101(path string) string {
	var stack []string
	slices := strings.Split(path, "/")
	for _, s := range slices {
		if s == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if s != "." && s != "" {
			stack = append(stack, s)
		}
	}
	return "/" + strings.Join(stack, "/")
}
