package review_20221226

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
		if node.min > val {
			node.min = val
		}
	}
}

func (t *trie) getMaxWithLimit(val int, limit int) int {
	node := t
	if node.min > limit {
		return -1
	}
	res := 0
	for i := L - 1; i >= 0; i-- {
		bit := val >> i & 1
		if node.children[bit^1] != nil && node.children[bit^1].min <= limit {
			res |= 1 << i
			bit ^= 1
		}
		node = node.children[bit]
	}
	return res
}

func maximizeXor(nums []int, queries [][]int) []int {
	res := make([]int, len(queries))
	t := &trie{
		min: math.MaxInt32,
	}
	for _, num := range nums {
		t.insert(num)
	}
	for i, query := range queries {
		res[i] = t.getMaxWithLimit(query[0], query[1])
	}
	return res
}
