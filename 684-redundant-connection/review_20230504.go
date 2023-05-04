package main

func findRedundantConnection_20230504(edges [][]int) []int {
	n := len(edges)
	parents := make([]int, n+1)
	for i := 0; i < len(parents); i++ {
		parents[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		// if
		if parents[x] != x {
			// find(parents[x])
			parents[x] = find(parents[x])
		}
		return parents[x]
	}
	var union func(from, to int) bool
	union = func(from, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return false
		}
		// notice
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
