package main

// https://www.youtube.com/watch?v=4hJ721ce010
func findRedundantConnection_20230708_2(edges [][]int) []int {
	parents := make([]int, len(edges)+1)
	sizes := make([]int, len(edges)+1)
	for i := range sizes {
		sizes[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		for parents[x] != x {
			parents[x] = parents[parents[x]]
			x = parents[x]
		}
		return x
	}

	for _, edge := range edges {
		u := edge[0]
		v := edge[1]

		// 没有初始化，初始化成自己
		if parents[u] == 0 {
			parents[u] = u
		}
		if parents[v] == 0 {
			parents[v] = v
		}

		// 找到根节点
		pu := find(u)
		pv := find(v)
		// 如果连通则返回这条边
		if pu == pv {
			return edge
		}

		// 如果不连通
		// 如果 pv 的叶子比 pu 大，则交换，把小的合并到大的
		if sizes[pu] < sizes[pv] {
			pu, pv = pv, pu
		}

		parents[pv] = pu
		sizes[pu] += sizes[pv]
	}
	return nil
}
