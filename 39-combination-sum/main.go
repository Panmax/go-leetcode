package main

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int

	var combs []int
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			// copyCombs := make([]int, len(combs))
			// for i := range combs {
			// 	copyCombs[i] = combs[i]
			// }

			// 需要把 combs 复制一份
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
