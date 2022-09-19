package main

func findAnagrams_20220919(s string, p string) []int {
	res := make([]int, 0)

	sLen := len(s)
	pLen := len(p)
	if sLen < pLen {
		return res
	}

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
