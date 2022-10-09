package main

import "sort"

func maxProfitAssignment_20221009(difficulty []int, profit []int, worker []int) int {
	var pairs Pairs
	for i := range difficulty {
		pairs = append(pairs, Pair{
			Diff:   difficulty[i],
			Profit: profit[i],
		})
	}
	sort.Sort(pairs)
	sort.Ints(worker)

	var res, best, maxIndex int
	for i := range worker {
		for maxIndex < len(pairs) && pairs[maxIndex].Diff <= worker[i] {
			best = Max(pairs[maxIndex].Profit, best)
			maxIndex++
		}
		res += best
	}
	return res
}
