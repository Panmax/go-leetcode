package main

func longestCommonPrefix_20240407(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp_20240407(prefix, strs[i])
		if prefix == "" {
			break
		}
	}
	return prefix
}

func lcp_20240407(prefix string, s string) string {
	n := min(len(prefix), len(s))
	i := 0
	for i < n && prefix[i] == s[i] {
		i++
	}
	return prefix[:i]
}
