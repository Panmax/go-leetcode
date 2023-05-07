package main

func combinationSum_20230507(candidates []int, target int) [][]int {
	var combs []int
	var res [][]int
	var dfs func(target int, idx int)
	dfs = func(target int, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			res = append(res, append([]int{}, combs...))
			return
		}
		dfs(target, idx+1)
		if target-candidates[idx] >= 0 {
			combs = append(combs, candidates[idx])
			dfs(target-candidates[idx], idx)
			combs = combs[:len(combs)-1]
		}
	}

	dfs(target, 0)

	return res
}
