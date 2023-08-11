package main

func findAnagrams_20230811(s string, p string) []int {
	sLen := len(s)
	pLen := len(p)
	if sLen < pLen {
		return nil
	}

	var res []int
	slotS := [26]int{}
	slotP := [26]int{}
	for i := 0; i < pLen; i++ {
		slotS[s[i]-'a']++
		slotP[p[i]-'a']++
	}
	if slotS == slotP {
		res = append(res, 0)
	}
	for i := pLen; i < sLen; i++ {
		slotS[s[i-pLen]-'a']--
		slotS[s[i]-'a']++
		if slotS == slotP {
			res = append(res, i-pLen+1)
		}
	}

	return res
}
