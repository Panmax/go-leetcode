package main

import "strconv"

func restoreIpAddresses_20221212(s string) []string {
	var res []string
	segments := make([]int, 4)
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
				res = append(res, ipAddr)
			}
			// 可以提前返回
			return
		}
		// 判断 start 而不是 id
		if start == len(s) {
			return
		}

		if s[start] == '0' {
			segments[id] = 0
			dfs(id+1, start+1)
		}

		var ip int
		for i := start; i < len(s); i++ {
			ip = ip*10 + int(s[i]-'0')
			if ip > 0 && ip <= 255 {
				segments[id] = ip
				dfs(id+1, i+1)
			} else {
				break
			}
		}
	}
	dfs(0, 0)
	return res
}
