package main

import "sort"

func inventoryManagement_20240109(stock []int, cnt int) []int {
	sort.Ints(stock)
	return stock[:cnt]
}
