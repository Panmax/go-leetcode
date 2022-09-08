package main

import "math"

const L = 30

type trie struct {
	children [2]*trie
	min      int
}

func (t *trie) insert(val int) {
	node := t
	if node.min > val {
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

func (t *trie) getMaxOr(val, limit int) int {
	node := t
	if node.min > limit {
		return -1
	}
	var ans int
	for i := L - 1; i >= 0; i-- {
		bit := val >> i & 1
		if node.children[bit^1] != nil && node.children[bit^1].min <= limit {
			ans |= 1 << i
			bit ^= 1
		}
		node = node.children[bit]
	}
	return ans
}

func maximizeXor(nums []int, queries [][]int) []int {
	t := &trie{min: math.MaxInt32}
	for _, num := range nums {
		t.insert(num)
	}
	ans := make([]int, len(queries))
	for i, query := range queries {
		ans[i] = t.getMaxOr(query[0], query[1])
	}
	return ans
}
