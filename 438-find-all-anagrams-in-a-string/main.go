package main

func findAnagrams(s string, p string) []int {
	sLen := len(s)
	pLen := len(p)
	if sLen < pLen {
		return []int{}
	}

	res := make([]int, 0)
	slotS := [26]int{}
	slotP := [26]int{}
	for i := range p {
		slotS[s[i]-'a']++
		slotP[p[i]-'a']++
	}
	if slotS == slotP {
		res = append(res, 0)
	}

	for i := pLen; i < sLen; i++ {
		slotS[s[i]-'a']++
		slotS[s[i-pLen]-'a']--
		if slotS == slotP {
			res = append(res, i-pLen+1)
		}
	}
	return res
}
