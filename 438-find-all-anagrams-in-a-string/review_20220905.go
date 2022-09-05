package main

func findAnagrams_20220905(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	res := make([]int, 0)
	slotS, slotP := [26]int{}, [26]int{}
	for i := range p {
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
			// 注意下边是 i-len(p)+1
			res = append(res, i-len(p)+1)
		}
	}

	return res
}
