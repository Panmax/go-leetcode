package main

import "sort"

func maxProfitAssignment_20230303(difficulty []int, profit []int, worker []int) int {
	var pairs Pairs
	for i, d := range difficulty {
		pairs = append(pairs, Pair{
			Diff:   d,
			Profit: profit[i],
		})
	}
	sort.Sort(pairs)
	sort.Ints(worker)
	ans := 0
	maxProfile := 0
	maxIndex := 0
	for _, w := range worker {
		for maxIndex < len(pairs) && pairs[maxIndex].Diff <= w {
			maxProfile = Max(maxProfile, pairs[maxIndex].Profit)
			maxIndex++
		}
		ans += maxProfile
	}

	return ans
}
