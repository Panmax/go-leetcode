package main

func rand10_20240425() int {
	for {
		n := (rand7()-1)*7 + (rand7() - 1)
		if n >= 1 && n <= 40 {
			return n%10 + 1
		}
	}
}
