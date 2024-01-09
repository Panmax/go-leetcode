package main

func combinationSum_20240109(candidates []int, target int) [][]int {
	var combs []int
	var ans [][]int
	var dfs func(target int, index int)
	dfs = func(target int, index int) {
		if index == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int{}, combs...)) // copy
			return                                       // return
		}
		dfs(target, index+1)
		if target-candidates[index] >= 0 {
			combs = append(combs, candidates[index])
			dfs(target-candidates[index], index)
			combs = combs[:len(combs)-1]
		}
	}

	dfs(target, 0)
	return ans
}
