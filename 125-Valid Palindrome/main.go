package main

import (
	"unicode"
)

func isPalindrome(s string) bool {
	p1, p2 := 0, len(s)-1

	for {
		if p1 == p2 || p2 < p1 {
			return true
		}

		r1 := rune(s[p1])
		r2 := rune(s[p2])

		l1 := unicode.ToLower(r1)
		l2 := unicode.ToLower(r2)

		if !isAlphaNumeric(l1) {
			p1++
			continue
		}

		if !isAlphaNumeric(l2) {
			p2--
			continue
		}

		if l1 != l2 {
			return false
		} else {
			p1++
			p2--
		}
	}
}

func isAlphaNumeric(r rune) bool {
	isLetter := (97 <= int(r)) && (int(r) <= 122)
	isNumber := (48 <= int(r)) && (int(r) <= 57)

	if isLetter || isNumber {
		return true
	}

	return false
}
