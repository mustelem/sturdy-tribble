package main

import (
	"strings"
)

func reverseWords(s string) string {
	var retVal string
	var words []string = strings.Split(s, " ")

	for i := len(words) - 1; i >= 0; i-- {
		if words[i] != "" {
			retVal += words[i]
			retVal += " "
		}
	}

	return strings.TrimSpace(retVal)
}
