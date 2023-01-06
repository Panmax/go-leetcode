package main

func findAnagrams_20230106(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	var res []int
	slotS := [26]int{}
	slotP := [26]int{}
	for i := 0; i < len(p); i++ {
		slotS[s[i]-'a']++
		slotP[p[i]-'a']++
	}
	if slotS == slotP {
		res = append(res, 0)
	}
	for i := len(p); i < len(s); i++ {
		slotS[s[i-len(p)]-'a']--
		slotS[s[i]-'a']++
		if slotS == slotP {
			// i-len(p)+1
			res = append(res, i-len(p)+1)
		}
	}
	return res
}
