package main

func canFinish_20231114(numCourses int, prerequisites [][]int) bool {
	var res []int
	edges := make(map[int][]int)
	indegs := make([]int, numCourses)
	for _, prerequisite := range prerequisites {
		edges[prerequisite[1]] = append(edges[prerequisite[1]], prerequisite[0])
		indegs[prerequisite[0]]++
	}
	var queue []int
	for i, indeg := range indegs {
		if indeg == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		res = append(res, i)
		for _, c := range edges[i] {
			indegs[c]--
			if indegs[c] == 0 {
				queue = append(queue, c)
			}
		}
	}
	return len(res) == numCourses
}
