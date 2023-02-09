package main

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	// 先判断是否有解
	m := make(map[string]bool, len(wordDict))
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
			// 倒序，不能直接改 path
			var reversePath []string
			for i := len(path) - 1; i >= 0; i-- {
				reversePath = append(reversePath, path[i])
			}
			res = append(res, strings.Join(reversePath, " "))
			return
		}
		for i := index - 1; i >= 0; i-- {
			suffix := s[i:index]
			if m[suffix] && dp[i] {
				path = append(path, suffix)
				dfs(i)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(len(s))
	return res
}

func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}
