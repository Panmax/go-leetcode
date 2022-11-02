package main

import "strings"

func simplifyPath_20221017(path string) string {
	var stack []string
	for _, name := range strings.Split(path, "/") {
		if name == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if name != "" && name != "." {
			stack = append(stack, name)
		}
	}
	res := strings.Join(stack, "/")
	return "/" + res
}

func main() {
	println(simplifyPath_20221017("/home/"))
	println(simplifyPath_20221017("/../"))
	println(simplifyPath_20221017("/a/./b/../../c/"))
}
