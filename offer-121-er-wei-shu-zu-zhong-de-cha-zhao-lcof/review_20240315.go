package main

func findTargetIn2DPlants_20240315(plants [][]int, target int) bool {
	if len(plants) == 0 {
		return false
	}
	m, n := len(plants), len(plants[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		num := plants[i][j]
		if num == target {
			return true
		} else if num < target {
			i++
		} else {
			j--
		}
	}
	return false
}
