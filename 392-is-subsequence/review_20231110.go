package main

func isSubsequence_20231110(s, t string) bool {
	var sPtr, tPtr int
	for sPtr < len(s) && tPtr < len(t) {
		if s[sPtr] == t[tPtr] {
			sPtr++
		}
		tPtr++
	}
	return sPtr == len(s)
}
