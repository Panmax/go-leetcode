package main

import "strconv"

func restoreIpAddresses_20230219(s string) []string {
	var res []string
	segments := make([]int, 4)
	var dfs func(id, index int)
	dfs = func(id, index int) {
		if id == 4 {
			if index == len(s) {
				ip := ""
				for i, segment := range segments {
					ip += strconv.Itoa(segment)
					if i < 3 {
						ip += "."
					}
				}
				res = append(res, ip)
			}
			return
		}
		if index == len(s) {
			return
		}
		if s[index] == '0' {
			segments[id] = 0
			dfs(id+1, index+1)
		}
		var segment int
		for i := index; i < len(s); i++ {
			segment = segment*10 + int(s[i]-'0')
			if segment > 0 && segment <= 255 {
				segments[id] = segment
				dfs(id+1, i+1)
			} else {
				break
			}
		}
	}
	dfs(0, 0)
	return res
}
