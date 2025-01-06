package main

import (
	"fmt"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	var result []string = make([]string, 300)
	resPtr := 0
	start := 0
	var lastline bool

	for {
		// intake(words []string, maxWidth int) string
		remaining := maxWidth
		payload := ""
		for i := start; i < len(words); i++ {
			start = i
			var sizeNextPart int
			if len(payload) > 0 {
				sizeNextPart = 1 + len(words[i])
			} else {
				sizeNextPart = len(words[i])
			}
			if sizeNextPart <= remaining {
				if remaining < maxWidth {
					payload += " "
					remaining -= 1
				}
				payload += words[i]
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
		var resul string
		// fmt.Println("payload:", payload)
		resul = justifyLine(payload, maxWidth, lastline)
		fmt.Println("result :", "\""+resul+"\"")

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
	extraSpace := i - len(s)
	var result string

	words := strings.Split(s, " ")
	places := len(words) - 1

	var spacers []string = make([]string, 100)

	for i := 0; i < places; i++ {
		spacers[i] = " "
	}

	if places < 1 {
		places = 1
	}

	for i := 0; i < extraSpace; i++ {
		spacers[i%places] += " "
	}

	result += words[0]
	for i := 0; i < places; i++ {
		result += spacers[i]
		if len(words) > 1 {
			result += words[i+1]
		}
	}

	if lastLine {
		spacer := ""
		for i := 0; i < extraSpace; i++ {
			spacer += " "
		}
		return s + spacer
	}

	return result
}

func main() {
	words := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	fullJustify(words, 20)
}
