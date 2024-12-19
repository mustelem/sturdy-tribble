package main

import (
	"bytes"
	"strings"
)

func reverseWords(s string) string {
	var words []string = strings.Split(s, " ")
	var buffer bytes.Buffer

	for i := len(words) - 1; i >= 0; i-- {
		if words[i] != "" {
			buffer.WriteString(words[i])
			buffer.WriteString(" ")
		}
	}

	return strings.TrimSpace(buffer.String())
}
