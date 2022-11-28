package main

import "strings"

func numberOfBeams(bank []string) int {
	var ret, last int
	for _, s := range bank {
		count := strings.Count(s, "1")
		if count <= 0 {
			continue
		}
		ret += last * count
		last = count
	}

	return ret
}
