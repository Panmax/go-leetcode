package main

func minNumberOfFrogs_20230115(croakOfFrogs string) int {
	var c, r, o, a, k int
	var ans int
	for _, char := range croakOfFrogs {
		if char == 'c' {
			if k > 0 {
				k--
			} else {
				ans++
			}
			c++
		} else if char == 'r' {
			c--
			r++
		} else if char == 'o' {
			r--
			o++
		} else if char == 'a' {
			o--
			a++
		} else if char == 'k' {
			a--
			k++
		}
		if c < 0 || r < 0 || o < 0 || a < 0 {
			break
		}
	}
	if c != 0 || r != 0 || o != 0 || a != 0 {
		return -1
	}
	return ans
}
