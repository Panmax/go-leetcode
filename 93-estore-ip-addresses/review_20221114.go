package main

import "strconv"

func restoreIpAddresses_20221114(s string) []string {
	segments := make([]int, 4)
	var ans []string
	var dfs func(id, start int)
	dfs = func(id, start int) {
		if id == 4 {
			if start == len(s) {
				ipAddr := ""
				for i, segment := range segments {
					ipAddr += strconv.Itoa(segment)
					if i < 3 {
						ipAddr += "."
					}
				}
				ans = append(ans, ipAddr)
			}
			return
		}
		if start == len(s) {
			return
		}
		if s[start] == '0' {
			segments[id] = 0
			dfs(id+1, start+1)
		}
		addr := 0
		for i := start; i < len(s); i++ {
			addr = addr*10 + int(s[i]-'0')
			if addr > 0 && addr <= 255 {
				segments[id] = addr
				dfs(id+1, i+1)
			} else {
				break
			}
		}
	}
	dfs(0, 0)

	return ans
}
