package main

import "strings"

func numberOfBeams_20230724(bank []string) int {
	res, last := 0, 0
	for _, s := range bank {
		count := strings.Count(s, "1")
		if count <= 0 {
			continue
		}
		res += last * count
		last = count
	}
	return res
}
