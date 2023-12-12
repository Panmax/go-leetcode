package main

import "sort"

func inventoryManagement_20231212(stock []int, cnt int) []int {
	sort.Ints(stock)
	return stock[:cnt]
}
