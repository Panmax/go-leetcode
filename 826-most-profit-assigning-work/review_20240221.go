package main

import "sort"

func maxProfitAssignment_20240221(difficulty []int, profit []int, worker []int) int {
	pairs := Pairs{}
	for i := range difficulty {
		pairs = append(pairs, Pair{
			Diff:   difficulty[i],
			Profit: profit[i],
		})
	}
	sort.Sort(pairs)
	sort.Ints(worker)
	j := 0
	best := 0
	ans := 0
	for _, w := range worker {
		for j < len(pairs) && w >= pairs[j].Diff {
			best = Max(pairs[j].Profit, best)
			j++
		}
		ans += best
	}
	return ans
}
