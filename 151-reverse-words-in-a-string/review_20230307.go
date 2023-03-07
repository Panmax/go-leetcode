package main

import "strings"

func reverseWords_20230307(s string) string {
	s = strings.TrimSpace(s)
	splits := strings.Split(s, " ")
	var nospaceSplits []string
	for _, split := range splits {
		if split != "" {
			nospaceSplits = append(nospaceSplits, split)
		}
	}
	for i := 0; i < len(nospaceSplits)/2; i++ {
		nospaceSplits[i], nospaceSplits[len(nospaceSplits)-i-1] = nospaceSplits[len(nospaceSplits)-i-1], nospaceSplits[i]
	}
	return strings.Join(nospaceSplits, " ")
}
