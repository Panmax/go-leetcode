package main

func findAnagrams_20220908(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	ans := make([]int, 0)
	slotS, slotP := [26]int{}, [26]int{}

	for i := range p {
		slotS[s[i]-'a']++
		slotP[p[i]-'a']++
	}
	if slotS == slotP {
		ans = append(ans, 0)
	}

	for i := len(p); i < len(s); i++ {
		slotS[s[i]-'a']++
		slotS[s[i-len(p)]-'a']--
		if slotS == slotP {
			ans = append(ans, i-len(p)+1)
		}
	}

	return ans
}
