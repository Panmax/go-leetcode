package main

import "sort"

func maxProfitAssignment_20220922(difficulty []int, profit []int, worker []int) int {
	var pairs Paris
	for i := range difficulty {
		pairs = append(pairs, Pair{
			Diff:   difficulty[i],
			Profit: profit[i],
		})
	}
	sort.Sort(pairs)
	sort.Ints(worker)

	var res int
	var bestIndex, bestProfit int

	for i := range worker {
		for bestIndex < len(pairs) && worker[i] >= pairs[bestIndex].Diff {

			if pairs[bestIndex].Profit > bestProfit {
				bestProfit = pairs[bestIndex].Profit
			}
			// 这个++ 放到判断下边
			bestIndex++
		}
		res += bestProfit
	}

	return res
}
