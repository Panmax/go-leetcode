package main

import (
	"strconv"
	"strings"
)

func exclusiveTime_20221001(n int, logs []string) []int {
	var stack []pair
	res := make([]int, n)
	for _, log := range logs {
		slice := strings.Split(log, ":")
		idx, _ := strconv.Atoi(slice[0])
		timestamp, _ := strconv.Atoi(slice[2])
		if slice[1] == "start" {
			if len(stack) > 0 {
				// 修改的是最后一个节点的 idx，而不是当前 idx
				res[stack[len(stack)-1].idx] += timestamp - stack[len(stack)-1].timestamp
				stack[len(stack)-1].timestamp = timestamp
			}
			stack = append(stack, pair{
				idx:       idx,
				timestamp: timestamp,
			})
		} else {
			// 需要记住最后一个节点后弹出
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[idx] += timestamp - p.timestamp + 1
			// 判空
			if len(stack) > 0 {
				stack[len(stack)-1].timestamp = timestamp + 1
			}
		}
	}
	return res
}
