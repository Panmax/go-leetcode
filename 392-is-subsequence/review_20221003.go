package main

func isSubsequence_20221003(s, t string) bool {
	sPtr, tPtr := 0, 0
	for sPtr < len(s) && tPtr < len(t) {
		if s[sPtr] == t[tPtr] {
			sPtr++
		}
		tPtr++
	}
	return sPtr == len(s)
}
