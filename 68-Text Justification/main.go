package main

import (
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	var result []string = make([]string, 150)
	var resPtr int = 0
	var start int = 0
	var lastline bool

	for {
		// intake(words []string, maxWidth int) string
		var remaining int = maxWidth
		var payload strings.Builder

		for i := start; i < len(words); i++ {
			start = i
			var sizeNextPart int
			if len(payload.String()) > 0 {
				sizeNextPart = 1 + len(words[i])
			} else {
				sizeNextPart = len(words[i])
			}
			if sizeNextPart <= remaining {
				if remaining < maxWidth {
					payload.WriteString(" ")
					remaining -= 1
				}
				payload.WriteString(words[i])
				remaining -= len(words[i])
				start++
			} else {
				break
			}
		}

		// justify line
		if start == len(words) {
			lastline = true
		}

		var resul string = justifyLine(payload.String(), maxWidth, lastline)

		//add to result array here
		result[resPtr] = resul
		resPtr++

		if start >= len(words) {
			break
		}
	}

	return result[0:resPtr]
}

func justifyLine(s string, i int, lastLine bool) string {
	var extraSpace int = i - len(s)
	var words []string = strings.Split(s, " ")
	var places int = len(words) - 1
	var spacers []strings.Builder = make([]strings.Builder, 100)

	for i := 0; i < places; i++ {
		spacers[i].WriteString(" ")
	}

	if places < 1 {
		places = 1
	}

	for i := 0; i < extraSpace; i++ {
		spacers[i%places].WriteString(" ")
	}

	var result strings.Builder
	result.WriteString(words[0])
	for i := 0; i < places; i++ {
		result.WriteString(spacers[i].String())
		if len(words) > 1 {
			result.WriteString(words[i+1])
		}
	}

	if lastLine {
		var spacer strings.Builder
		spacer.WriteString(s)
		for i := 0; i < extraSpace; i++ {
			spacer.WriteString(" ")
		}
		return spacer.String()
	}

	return result.String()
}
