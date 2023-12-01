package main

func longestCommonPrefix_20231201(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp_20231201(prefix, strs[i])
		if prefix == "" {
			break
		}
	}

	return prefix
}

func lcp_20231201(prefix string, str string) string {
	n := min(len(prefix), len(str))
	i := 0
	for i < n && prefix[i] == str[i] {
		i++
	}
	return prefix[:i]
}
