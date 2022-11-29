package main

import "strings"

func numberOfBeams_20221129(bank []string) int {
	var ret, last int
	for _, s := range bank {
		n := strings.Count(s, "1")
		if n <= 0 {
			continue
		}
		ret += last * n
		last = n
	}
	return ret
}
