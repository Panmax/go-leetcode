package main

func checkInclusion_20220818(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Slot := [26]byte{}
	s2SLot := [26]byte{}
	// 需要 - 'a'
	for i := 0; i < len(s1); i++ {
		s1Slot[s1[i]]++
		s2SLot[s2[i]]++
	}
	if s1Slot == s2SLot {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		s2SLot[s2[i]]++
		s2SLot[s2[i-len(s1)]]--
		if s1Slot == s2SLot {
			return true
		}
	}
	return false
}
