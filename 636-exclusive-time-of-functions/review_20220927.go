package main

import (
	"strconv"
	"strings"
)

func exclusiveTime_20220927(n int, logs []string) []int {
	res := make([]int, n)
	var stack []pair
	for _, log := range logs {
		slice := strings.Split(log, ":")
		idx, _ := strconv.Atoi(slice[0])
		timestamp, _ := strconv.Atoi(slice[2])
		if slice[1] == "start" {
			if len(stack) > 0 {
				res[stack[len(stack)-1].idx] += timestamp - stack[len(stack)-1].timestamp
				stack[len(stack)-1].timestamp = timestamp
			}

			stack = append(stack, pair{
				idx:       idx,
				timestamp: timestamp,
			})
		} else {
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[idx] += timestamp - p.timestamp + 1
			// 判断大于0
			if len(stack) > 0 {
				stack[len(stack)-1].timestamp = timestamp + 1
			}

		}
	}
	return res
}
