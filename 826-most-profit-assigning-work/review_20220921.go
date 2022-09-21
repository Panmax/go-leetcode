package main

import "sort"

func maxProfitAssignment_20220921(difficulty []int, profit []int, worker []int) int {
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
	var bestIndex, bestProfit int
	for i := range worker {
		for bestIndex < len(pairs) && worker[i] >= pairs[bestIndex].Diff {
			bestProfit = Max(bestProfit, pairs[bestIndex].Profit)
			bestIndex++
		}
		ans += bestProfit
	}
	return ans
}
