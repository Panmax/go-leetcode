package main

// 在 strs 中寻找 s，这个 s 不是其他字符串的子序列
// 判断 a 是否为 b 的子序列 --> b 是否可以通过删除字符得到 a
func findLUSlength(strs []string) int {
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isSub(s1, s2 string) bool {
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
		}
		j++
	}
	return i == len(s1)
}
