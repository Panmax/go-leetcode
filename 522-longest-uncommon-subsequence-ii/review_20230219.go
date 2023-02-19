package main

func findLUSlength_20230219(strs []string) int {
	ans := -1
	for i := 0; i < len(strs); i++ {
		check := true
		for j := 0; j < len(strs); j++ {
			if i != j && isSub(strs[i], strs[j]) {
				check = false
				break
			}
		}
		if check {
			ans = max(ans, len(strs[i]))
		}
	}
	return ans
}
