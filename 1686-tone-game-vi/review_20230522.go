package main

import "sort"

func stoneGameVI_20230522(aliceValues []int, bobValues []int) int {
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
	var alice, bob int
	for i := 0; i < len(values); i++ {
		if i%2 == 0 {
			alice += aliceValues[values[i].index]
		} else {
			bob += bobValues[values[i].index]
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
