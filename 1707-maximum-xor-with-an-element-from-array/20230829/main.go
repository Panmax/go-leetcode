package main

import "math"

const L = 30

type trie struct {
	children [2]*trie
	min      int
}

func (t *trie) insert(num int) {
	node := t
	if node.min > num {
		node.min = num
	}
	for i := L - 1; i >= 0; i-- {
		bit := num >> i & 1
		if node.children[bit] == nil {
			node.children[bit] = &trie{min: num}
		}
		node = node.children[bit]
		if node.min > num {
			node.min = num
		}
	}
}

func (t *trie) getMaxWithLimit(val, limit int) int {
	node := t
	if node.min > limit {
		return -1
	}
	res := 0
	for i := L - 1; i >= 0; i-- {
		bit := val >> i & 1
		if node.children[bit^1] != nil && node.children[bit^1].min <= limit {
			bit ^= 1
			res |= 1 << i
		}
		node = node.children[bit]
	}
	return res
}

func maximizeXor(nums []int, queries [][]int) []int {
	t := &trie{min: math.MaxInt32}
	for _, num := range nums {
		t.insert(num)
	}
	ans := make([]int, len(queries))
	for i, query := range queries {
		ans[i] = t.getMaxWithLimit(query[0], query[1])
	}
	return ans
}
