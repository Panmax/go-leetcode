package main

func wordBreak_20230309(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, w := range wordDict {
		m[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && m[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
