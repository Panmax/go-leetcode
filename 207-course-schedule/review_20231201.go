package main

func canFinish_20231201(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
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
	var res []int
	for len(queue) > 0 {
		i := queue[0]
		res = append(res, 0)
		queue = queue[1:]
		for _, c := range edges[i] {
			indegs[c]--
			if indegs[c] == 0 {
				queue = append(queue, c)
			}
		}

	}
	return len(res) == numCourses
}
