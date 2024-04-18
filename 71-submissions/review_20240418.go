package main

import "strings"

func simplifyPath_20240418(path string) string {
	var stack []string
	paths := strings.Split(path, "/")
	for _, s := range paths {
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
