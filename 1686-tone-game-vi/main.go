package main

import "sort"

type Value struct {
	index int
	value int
}

func stoneGameVI(aliceValues []int, bobValues []int) int {
	n := len(aliceValues)
	var values []*Value
	for i := 0; i < n; i++ {
		values = append(values, &Value{
			index: i,
			value: aliceValues[i] + bobValues[i],
		})
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i].value > values[j].value
	})
	alice := 0
	bob := 0
	for i, value := range values {
		if i%2 == 0 {
			alice += aliceValues[value.index]
		} else {
			bob += bobValues[value.index]
		}
	}
	if alice > bob {
		return 1
	} else if alice < bob {
		return -1
	} else {
		return 0
	}
}
