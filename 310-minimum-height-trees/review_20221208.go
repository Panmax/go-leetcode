package main

func findMinHeightTrees_20221208(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	degree := make([]int, n)
	m := make(map[int][]int)
	for _, edge := range edges {
		m[edge[0]] = append(m[edge[0]], edge[1])
		m[edge[1]] = append(m[edge[1]], edge[0])
		degree[edge[0]]++
		degree[edge[1]]++
	}
	var queue []int
	for i, d := range degree {
		if d == 1 {
			queue = append(queue, i)
		}
	}
	var ans []int
	for len(queue) > 0 {
		var level []int
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			level = append(level, cur)
			for _, nei := range m[cur] {
				degree[nei]--
				if degree[nei] == 1 {
					queue = append(queue, nei)
				}
			}
		}

		queue = queue[size:]
		ans = level
	}
	return ans
}
