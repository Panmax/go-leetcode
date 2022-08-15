package main

import "math/rand"

func (s *Solution) GetRandom_20220815() int {
	var res int
	curr := s.head
	i := 1
	for curr != nil {
		if rand.Intn(i) == 0 {
			res = curr.Val
		}
		i++
		curr = curr.Next
	}

	return res
}
