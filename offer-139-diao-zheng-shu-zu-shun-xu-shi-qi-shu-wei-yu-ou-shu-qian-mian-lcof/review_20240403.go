package main

func trainingPlan_20240403(actions []int) []int {
	i, j := 0, len(actions)-1
	for i < j {
		for i < j {
			if actions[i]%2 == 0 {
				break
			}
			i++
		}
		for i < j {
			if actions[j]%2 == 1 {
				break
			}
			j--
		}
		if i < j {
			actions[i], actions[j] = actions[j], actions[i]
		}
		i++
		j--
	}
	return actions
}
