package main

import "sort"

func maxProfitAssignment_20220926(difficulty []int, profit []int, worker []int) int {
	var pairs Paris
	for i := range difficulty {
		pairs = append(pairs, Pair{
			Diff:   difficulty[i],
			Profit: profit[i],
		})
	}
	sort.Sort(pairs)
	sort.Ints(worker)

	var ans int
	var best int
	var maxIndex int
	for i := range worker {
		for maxIndex < len(pairs) && pairs[maxIndex].Diff <= worker[i] {
			best = Max(best, pairs[maxIndex].Profit)
			maxIndex++
		}
		ans += best
	}
	return ans
}
