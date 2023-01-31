package main

import "strings"

func numberOfBeams_20230131(bank []string) int {
	var ans, last int
	for _, s := range bank {
		count := strings.Count(s, "1")
		if count <= 0 {
			continue
		}
		ans += last * count
		last = count
	}
	return ans
}
