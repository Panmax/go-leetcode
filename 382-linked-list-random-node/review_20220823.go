package main

import "math/rand"

func (s *Solution) GetRandom_20220823() int {
	var res int
	curr := s.head
	i := 1
	for curr != nil {
		if rand.Intn(i) == 0 {
			res = curr.Val
		}
		curr = curr.Next
		i++
	}
	return res
}
