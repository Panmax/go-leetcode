package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	len1 := len(s1)
	len2 := len(s2)
	len3 := len(s3)
	if len1+len2 != len3 {
		return false
	}

	dp := make([][]bool, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]bool, len2+1)
	}
	dp[0][0] = true
	for i := 0; i <= len1; i++ {
		for j := 0; j <= len2; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}

	return dp[len1][len2]
}

func main() {
	fmt.Println(isInterleave("aabcc",
		"dbbca",
		"aadbbcbcac"))
}
