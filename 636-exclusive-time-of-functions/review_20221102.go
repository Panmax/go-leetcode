package main

import (
	"strconv"
	"strings"
)

func exclusiveTime_20221102(n int, logs []string) []int {
	res := make([]int, n)
	var stack []pair
	for _, log := range logs {
		slices := strings.Split(log, ":")
		idx, _ := strconv.Atoi(slices[0])
		timestamp, _ := strconv.Atoi(slices[2])
		if slices[1] == "start" {
			if len(stack) > 0 {
				res[stack[len(stack)-1].idx] += timestamp - stack[len(stack)-1].timestamp
				stack[len(stack)-1].timestamp = timestamp
			}
			stack = append(stack, pair{idx: idx, timestamp: timestamp})
		} else {
			p := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[p.idx] += timestamp - p.timestamp + 1

			if len(stack) > 0 {
				stack[len(stack)-1].timestamp = timestamp + 1
			}
		}
	}

	return res
}
