package main

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	degree := make([]int, n)
	m := make(map[int][]int)
	for _, pair := range edges {
		m[pair[0]] = append(m[pair[0]], pair[1])
		m[pair[1]] = append(m[pair[1]], pair[0])
		degree[pair[0]]++
		degree[pair[1]]++
	}
	queue := make([]int, 0)
	for i, d := range degree {
		if d == 1 {
			queue = append(queue, i)
		}
	}

	var ans []int
	for len(queue) > 0 {
		level := make([]int, 0)
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
