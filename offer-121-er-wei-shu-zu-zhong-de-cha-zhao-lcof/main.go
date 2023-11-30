package main

func findTargetIn2DPlants(plants [][]int, target int) bool {
	if len(plants) == 0 {
		return false
	}
	m, n := len(plants), len(plants[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		if plants[i][j] == target {
			return true
		} else if plants[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
