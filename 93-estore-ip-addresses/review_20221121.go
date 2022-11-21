package main

import "strconv"

func restoreIpAddresses_20221121(s string) []string {
	segments := make([]int, 4)
	res := make([]string, 0)
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
			return
		}

		if start == len(s) {
			return
		}

		if s[start] == '0' {
			segments[id] = 0
			dfs(id+1, start+1)
		}

		ip := 0
		// i := start
		for i := start; i < len(s); i++ {
			ip = ip*10 + int(s[i]-'0')
			if ip > 0 && ip <= 255 {
				segments[id] = ip
				// i + 1
				dfs(id+1, i+1)
			} else {
				break
			}
		}
	}

	dfs(0, 0)
	return res
}
