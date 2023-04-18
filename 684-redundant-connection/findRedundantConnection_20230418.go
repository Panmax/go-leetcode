package main

func findRedundantConnection_20230418(edges [][]int) []int {
	n := len(edges)
	parents := make([]int, n+1)
	for i := range parents {
		parents[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parents[x] != x {
			// find(parents[x])
			parents[x] = find(parents[x])
		}
		return parents[x]
	}
	var union func(int, int) bool
	union = func(from int, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return false
		}
		parents[x] = y
		return true
	}
	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return edge
		}
	}
	return nil
}
