package main

func findAnagrams_20221019(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	slotS := [26]int{}
	slotP := [26]int{}
	for i := 0; i < len(p); i++ {
		slotS[s[i]-'a']++
		slotP[p[i]-'a']++
	}
	var res []int
	if slotS == slotP {
		res = append(res, 0)
	}
	for i := len(p); i < len(s); i++ {
		slotS[s[i]-'a']++
		slotS[s[i-len(p)]-'a']--
		if slotS == slotP {
			res = append(res, i-len(p)+1)
		}
	}

	return res
}
