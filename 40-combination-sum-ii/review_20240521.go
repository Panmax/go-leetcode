package main

import "sort"

func combinationSum2_20240521(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var freq [][2]int
	for _, candidate := range candidates {
		if freq == nil || candidate != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{candidate, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	var combs []int
	var ans [][]int
	var dfs func(target int, index int)
	dfs = func(target int, index int) {
		if target == 0 {
			ans = append(ans, append([]int{}, combs...))
			return
		}
		if index == len(freq) || target < 0 || target < freq[index][0] {
			return
		}
		dfs(target, index+1)
		for i := 1; i <= freq[index][1]; i++ {
			combs = append(combs, freq[index][0])
			dfs(target-i*freq[index][0], index+1)
		}
		combs = combs[:len(combs)-freq[index][1]]
	}
	dfs(target, 0)
	return ans
}
