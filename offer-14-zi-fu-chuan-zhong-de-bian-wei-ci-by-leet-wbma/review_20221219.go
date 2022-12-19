package main

func checkInclusion_20221219(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1 > l2 {
		return false
	}
	slot1, slot2 := [26]int{}, [26]int{}
	for i := 0; i < l1; i++ {
		slot1[s1[i]-'a']++
		slot2[s2[i]-'a']++
	}
	if slot1 == slot2 {
		return true
	}
	for i := l1; i < l2; i++ {
		slot2[s2[i-l1]-'a']--
		slot2[s2[i]-'a']++
		if slot1 == slot2 {
			return true
		}
	}

	return false
}
