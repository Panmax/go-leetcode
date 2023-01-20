package main

func minNumberOfFrogs_20230118(croakOfFrogs string) int {
	var ans int
	var c, r, o, a, k int
	for i := 0; i < len(croakOfFrogs); i++ {
		if croakOfFrogs[i] == 'c' {
			if k > 0 {
				k--
			} else {
				ans++
			}
			c++
		} else if croakOfFrogs[i] == 'r' {
			c--
			r++
		} else if croakOfFrogs[i] == 'o' {
			r--
			o++
		} else if croakOfFrogs[i] == 'a' {
			o--
			a++
		} else if croakOfFrogs[i] == 'k' {
			a--
			k++
		}
		if c < 0 || r < 0 || o < 0 || a < 0 {
			return -1
		}
	}
	if c != 0 || r != 0 || o != 0 || a != 0 {
		return -1
	}
	return ans
}
