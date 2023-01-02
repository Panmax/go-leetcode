package main

func isIsomorphic_20230102(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)
	for i := range s {
		x := s[i]
		y := t[i]
		if m1[x] > 0 || m2[y] > 0 {
			if m1[x] != y || m2[y] != x {
				return false
			}
			continue
		}
		m1[x] = y
		m2[y] = x
	}
	return true
}
