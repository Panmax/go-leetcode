package main

func combinationSum_20230628(candidates []int, target int) [][]int {
	var combs []int
	var ans [][]int
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), combs...))
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
	return ans
}
