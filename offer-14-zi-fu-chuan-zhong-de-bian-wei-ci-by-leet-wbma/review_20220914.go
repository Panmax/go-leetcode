package main

func checkInclusion_20220914(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l2 < l1 {
		return false
	}
	slotS1, slotS2 := [26]int{}, [26]int{}
	for i := range s1 {
		slotS1[s1[i]-'a']++
		slotS2[s2[i]-'a']++
	}
	if slotS1 == slotS2 {
		return true
	}
	for i := l1; i < l2; i++ {
		slotS2[s2[i-l1]-'a']--
		slotS2[s2[i]-'a']++
		if slotS1 == slotS2 {
			return true
		}
	}
	return false
}
