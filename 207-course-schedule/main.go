package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	res := make([]int, 0, numCourses)
	edge := make(map[int][]int)
	indeg := make([]int, numCourses)
	for _, prerequisite := range prerequisites {
		edge[prerequisite[1]] = append(edge[prerequisite[1]], prerequisite[0])
		indeg[prerequisite[0]]++
	}
	q := make([]int, 0)
	for i := range indeg {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		res = append(res, c)
		for _, i := range edge[c] {
			indeg[i]--
			if indeg[i] == 0 {
				q = append(q, i)
			}
		}
	}
	return len(res) == numCourses
}
