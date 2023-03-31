package main

import "sort"

func stoneGameVI_20230331(aliceValues []int, bobValues []int) int {
	var values []*Value
	for i := range aliceValues {
		values = append(values, &Value{
			index: i,
			value: aliceValues[i] + bobValues[i],
		})
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i].value > values[j].value
	})
	var alice int
	var bob int
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
