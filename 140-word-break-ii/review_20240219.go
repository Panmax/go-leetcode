package main

import "strings"

func wordBreak_20240219(s string, wordDict []string) []string {
	m := make(map[string]bool)
	for _, word := range wordDict {
		m[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && m[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	if !dp[len(s)] {
		return nil
	}

	var res []string
	var path []string
	var dfs func(index int)
	dfs = func(index int) {
		if index == 0 {
			var reversePath []string
			for i := len(path) - 1; i >= 0; i-- {
				reversePath = append(reversePath, path[i])
			}
			res = append(res, strings.Join(reversePath, " "))
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
