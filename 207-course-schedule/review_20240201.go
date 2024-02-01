package main

func canFinish_20240201(numCourses int, prerequisites [][]int) bool {
	var res []int
	edges := make([][]int, numCourses)
	indes := make([]int, numCourses)
	for _, prerequisite := range prerequisites {
		edges[prerequisite[1]] = append(edges[prerequisite[1]], prerequisite[0])
		indes[prerequisite[0]]++
	}
	var queue []int
	for i, ind := range indes {
		if ind == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		q := queue[0]
		res = append(res, q)
		queue = queue[1:]
		for _, c := range edges[q] {
			indes[c]--
			if indes[c] == 0 {
				queue = append(queue, c)
			}
		}
	}
	return len(res) == numCourses
}
