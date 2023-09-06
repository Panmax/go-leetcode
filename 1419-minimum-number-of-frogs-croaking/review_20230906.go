package main

func minNumberOfFrogs_20230906(croakOfFrogs string) int {
	var ans int
	var c, r, o, a, k int
	for _, char := range croakOfFrogs {
		switch char {
		case 'c':
			if k > 0 {
				k--
			} else {
				ans++
			}
			c++
		case 'r':
			c--
			r++
		case 'o':
			r--
			o++
		case 'a':
			o--
			a++
		case 'k':
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
