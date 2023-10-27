package main

func longestCommonPrefix_20231027(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp_20231027(prefix, strs[i])
		if prefix == "" {
			break
		}
	}
	return prefix
}

func lcp_20231027(prefix string, s string) string {
	n := min(len(prefix), len(s))
	var i int
	for i < n && prefix[i] == s[i] {
		i++
	}
	return prefix[:i]
}
