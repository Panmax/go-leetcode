package main

import "math"

// 定义 L
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
		// & 1
		bit := val >> i & 1
		if node.children[bit] == nil {
			node.children[bit] = &trie{min: val}

		}
		// 放到外边
		node = node.children[bit]
		if val < node.min {
			node.min = val
		}
	}
}

func (t *trie) getMaxXorWithLimit(val int, limit int) int {
	node := t
	if limit < node.min {
		return -1
	}
	ans := 0
	for i := L - 1; i >= 0; i-- {
		bit := val >> i & 1
		// node.children[bit^1].min <=limit
		if node.children[bit^1] != nil && node.children[bit^1].min <= limit {
			ans |= 1 << i
			bit ^= 1
		}
		node = node.children[bit]
	}
	return ans
}

func maximizeXor(nums []int, queries [][]int) []int {
	ans := make([]int, len(queries))

	t := &trie{min: math.MaxInt32}
	for _, num := range nums {
		t.insert(num)
	}
	for i, query := range queries {
		ans[i] = t.getMaxXorWithLimit(query[0], query[1])
	}

	return ans
}
