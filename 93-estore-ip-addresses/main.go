package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	segments := make([]int, 4)
	var res []string
	var dfs func(segNum, start int)
	dfs = func(segNum, start int) {
		if segNum == 4 {
			if start == len(s) {
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

		if start == len(s) {
			return
		}

		if s[start] == '0' {
			segments[segNum] = 0
			dfs(segNum+1, start+1)
		}

		addr := 0
		for i := start; i < len(s); i++ {
			addr = addr*10 + int(s[i]-'0')
			if addr > 0 && addr <= 255 {
				segments[segNum] = addr
				// i+1
				dfs(segNum+1, i+1)
			} else {
				break
			}
		}
	}
	dfs(0, 0)
	return res
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}
