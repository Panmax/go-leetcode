package main

func findRedundantConnection(edges [][]int) []int {
	parents := make([]int, len(edges)+1)
	for i := range parents {
		parents[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parents[x] != x {
			parents[x] = find(parents[x])
		}
		return parents[x]
	}
	union := func(from, to int) bool {
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
