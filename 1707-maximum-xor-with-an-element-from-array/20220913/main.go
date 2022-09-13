package main

import "math"

const L = 30

type trie struct {
	children [2]*trie
	min      int
}

func (t *trie) insert(val int) {
	node := t
	if val < node.min {
		node.min = val
	}
	for i := L - 1; i >= 0; i-- {
		bit := val >> i & 1
		if node.children[bit] == nil {
			node.children[bit] = &trie{min: val}
		}
		node = node.children[bit]
		if val < node.min {
			node.min = val
		}
	}
}

func (t *trie) getMax(x, m int) int {
	node := t
	if node.min > m {
		return -1
	}
	res := 0
	for i := L - 1; i >= 0; i-- {
		bit := x >> i & 1
		if node.children[bit^1] != nil && node.children[bit^1].min <= m {
			res |= 1 << i
			bit = bit ^ 1
		}
		node = node.children[bit]
	}
	return res
}

func maximizeXor(nums []int, queries [][]int) []int {
	// 初始值 MaxInt32
	t := &trie{min: math.MaxInt32}
	for _, num := range nums {
		t.insert(num)
	}
	res := make([]int, len(queries))
	for i, query := range queries {
		res[i] = t.getMax(query[0], query[1])
	}
	return res
}
