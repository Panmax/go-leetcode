package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp(prefix, strs[i])
		if prefix == "" {
			break
		}
	}
	return prefix
}

func lcp(prefix string, s string) string {
	n := min(len(prefix), len(s))
	i := 0
	for i < n && prefix[i] == s[i] {
		i++
	}
	return prefix[:i]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
