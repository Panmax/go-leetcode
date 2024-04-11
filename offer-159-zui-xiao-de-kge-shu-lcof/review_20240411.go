package main

import "sort"

func inventoryManagement_20240411(stock []int, cnt int) []int {
	sort.Ints(stock)
	return stock[:cnt]
}
