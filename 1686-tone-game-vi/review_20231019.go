package main

import "sort"

func stoneGameVI_20231019(aliceValues []int, bobValues []int) int {
	values := make([]*Value, len(aliceValues))
	for i := 0; i < len(aliceValues); i++ {
		values[i] = &Value{
			index: i,
			value: aliceValues[i] + bobValues[i],
		}
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
	} else if bob > alice {
		return -1
	} else {
		return 0
	}
}
