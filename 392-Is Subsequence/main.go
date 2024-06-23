package main

func isSubsequence(s string, t string) bool {
	lens := len(s)
	lent := len(t)

	if lens > lent {
		return false
	}

	if lens == 0 {
		return true
	}

	var p1 int

	for i := 0; i < lent; i++ {
		if s[p1] == t[i] {
			if p1 < lens-1 {
				p1++
			} else {
				return true
			}
		}
	}

	return p1 == lens
}
