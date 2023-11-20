package main

func checkInclusion_20231120(s1 string, s2 string) bool {
	n1, n2 := len(s1), len(s2)
	if n1 > n2 {
		return false
	}
	slot1, slot2 := [26]int{}, [26]int{}
	for i := 0; i < n1; i++ {
		slot1[s1[i]-'a']++
		slot2[s2[i]-'a']++
	}
	if slot1 == slot2 {
		return true
	}
	for i := n1; i < n2; i++ {
		slot2[s2[i-n1]-'a']--
		slot2[s2[i]-'a']++
		if slot1 == slot2 {
			return true
		}
	}
	return false
}
