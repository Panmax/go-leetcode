package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	// 构造二维 freq 数组，freq[i][0]标识值，freq[i][1]标示这个值出现次数
	sort.Ints(candidates)
	var freq [][2]int
	for _, num := range candidates {
		if freq == nil || num != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	// 下边使用构造出的 freq 进行计算，而不是原始 candidates
	var combs []int
	var ans [][]int
	var dfs func(target int, index int)
	dfs = func(target int, index int) {
		if target == 0 {
			ans = append(ans, append([]int{}, combs...)) // copy
			return                                       // return
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
