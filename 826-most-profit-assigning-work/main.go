package main

import "sort"

type Pair struct {
	Diff   int
	Profit int
}

type Paris []Pair

func (p Paris) Less(i, j int) bool {
	return p[i].Diff < p[j].Diff
}

func (p Paris) Len() int {
	return len(p)
}

func (p Paris) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
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
	// best 和 j 是精髓，可以少遍历一些
	var best int
	j := 0
	for i := range worker {
		for j < len(pairs) && worker[i] >= pairs[j].Diff {
			best = Max(pairs[j].Profit, best)
			j++
		}
		ans += best
	}

	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
