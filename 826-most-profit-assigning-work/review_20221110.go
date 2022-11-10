package main

import "sort"

func maxProfitAssignment_20221110(difficulty []int, profit []int, worker []int) int {
	var pairs Pairs
	for i, diff := range difficulty {
		pairs = append(pairs, Pair{
			Diff:   diff,
			Profit: profit[i],
		})
	}
	sort.Sort(pairs)
	sort.Ints(worker)

	ant := 0
	bestProfit := 0
	maxIndex := 0
	for _, w := range worker {
		for maxIndex < len(pairs) && pairs[maxIndex].Diff <= w {
			bestProfit = Max(bestProfit, pairs[maxIndex].Profit)
			maxIndex++
		}
		ant += bestProfit
	}

	return ant
}
