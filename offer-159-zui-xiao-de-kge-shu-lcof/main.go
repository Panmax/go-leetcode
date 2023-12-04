package main

import "sort"

func inventoryManagement(stock []int, cnt int) []int {
	sort.Ints(stock)
	return stock[:cnt]
}
