package main

func checkInclusion_20220804(s1 string, s2 string) bool {
	s1Len := len(s1)
	s2Len := len(s2)
	if s1Len > s2Len {
		return false
	}
	s1Slot := [26]int{}
	s2Slot := [26]int{}
	for i := range s1 {
		s1Slot[s1[i]-'a']++
		s2Slot[s2[i]-'a']++
	}
	if s1Slot == s2Slot {
		return true
	}
	for i := s1Len; i < s2Len; i++ {
		s2Slot[s2[i]-'a']++
		s2Slot[s2[i-s1Len]-'a']--
		if s1Slot == s2Slot {
			return true
		}
	}
	return false
}
