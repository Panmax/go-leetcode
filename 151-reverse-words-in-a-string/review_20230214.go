package main

import "strings"

func reverseWords_20230214(s string) string {
	s = strings.Trim(s, " ")
	slice := strings.Split(s, " ")
	// 去掉中间的空格
	var noSpaceSlice []string
	for _, val := range slice {
		if val != "" {
			noSpaceSlice = append(noSpaceSlice, val)
		}
	}
	for i := 0; i < len(noSpaceSlice)/2; i++ {
		noSpaceSlice[i], noSpaceSlice[len(noSpaceSlice)-i-1] = noSpaceSlice[len(noSpaceSlice)-i-1], noSpaceSlice[i]
	}
	s = strings.Join(noSpaceSlice, " ")
	return s
}

func main() {
	reverseWords_20230214("a good   example")
}
