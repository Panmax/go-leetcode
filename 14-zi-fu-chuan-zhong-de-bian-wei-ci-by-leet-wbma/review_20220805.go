package main

func checkInclusion_20220805(s1 string, s2 string) bool {
	s1Len := len(s1)
	s2Len := len(s2)
	if s1Len > s2Len {
		return false
	}
	slot1 := [26]int{}
	slot2 := [26]int{}

	for i := range s1 {
		slot1[s1[i]-'a']++
		slot2[s2[i]-'a']++
	}
	if slot1 == slot2 {
		return true
	}
	for i := s1Len; i < s2Len; i++ {
		slot2[s2[i]-'a']++
		slot2[s2[i-s1Len]-'a']--
		if slot1 == slot2 {
			return true
		}
	}
	return false
}
