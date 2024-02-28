package main

import (
	"strconv"
	"strings"
)

func compareVersion_20240228(version1 string, version2 string) int {
	v1s := strings.Split(version1, ".")
	v2s := strings.Split(version2, ".")
	for i := 0; i < len(v1s) || i < len(v2s); i++ {
		var x, y int
		if i < len(v1s) {
			x, _ = strconv.Atoi(v1s[i])
		}
		if i < len(v2s) {
			y, _ = strconv.Atoi(v2s[i])
		}
		if x > y {
			return 1
		} else if x < y {
			return -1
		}
	}
	return 0
}
