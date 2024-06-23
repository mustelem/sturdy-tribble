package main

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	if len(s) == 0 {
		return true
	}

	var p1 int

	for i := 0; i < len(t); i++ {
		if s[p1] == t[i] {
			if p1 < len(s)-1 {
				p1++
			} else {
				return true
			}
		}
	}

	return p1 == len(s)
}
