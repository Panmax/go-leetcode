package main

import "math/rand"

func (s *Solution) GetRandom_20220917() int {
	node := s.head
	var res int
	for i := 1; node != nil; i++ {
		if rand.Intn(i) == 0 {
			res = node.Val
		}
		node = node.Next
	}

	return res
}
