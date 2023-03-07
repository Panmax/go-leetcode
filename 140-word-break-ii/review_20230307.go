package main

import "strings"

func wordBreak_20230307(s string, wordDict []string) []string {
	m := make(map[string]bool)
	for _, w := range wordDict {
		m[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	n := len(s)
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && m[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	if !dp[n] {
		return nil
	}

	var path []string
	var res []string
	var dfs func(index int)
	dfs = func(index int) {
		if index == 0 {
			var reversedPath []string
			for i := len(path) - 1; i >= 0; i-- {
				reversedPath = append(reversedPath, path[i])
			}
			res = append(res, strings.Join(reversedPath, " "))
		}

		for i := index - 1; i >= 0; i-- {
			if dp[i] && m[s[i:index]] {
				path = append(path, s[i:index])
				dfs(i)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(len(s))
	return res
}
