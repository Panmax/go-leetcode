package main

func findRedundantConnection_20240301(edges [][]int) []int {
	parents := make([]int, len(edges)+1)
	for i := range parents {
		parents[i] = i
	}
	size := make([]int, len(edges)+1)
	for i := range size {
		size[i] = 1
	}

	var findParent func(int) int
	findParent = func(x int) int {
		for parents[x] != x {
			parents[x] = parents[parents[x]]
			x = parents[x]
		}
		return x
	}

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		pu := findParent(u)
		pv := findParent(v)
		if pu == pv {
			return edge
		}
		if size[pu] < size[pv] {
			pu, pv = pv, pu
		}
		parents[pu] = pv
		size[pu] += size[pv]
	}
	return nil
}
