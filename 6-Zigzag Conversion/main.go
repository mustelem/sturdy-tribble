package main

import (
	"bytes"
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	var buffer bytes.Buffer
	letters := strings.Split(s, "")

	i := 0
	width := numRows + max(0, numRows-2)
	diff := width

	for j := 0; j < numRows; j++ {
		i = j
		buffer.WriteString(letters[i])

		diff = width - max(0, 2*j)
		if diff == 0 {
			diff = width - diff
		}
		for i+diff < len(letters) {
			buffer.WriteString(letters[i+diff])
			i += diff

			diff = width - diff
			if diff == 0 {
				diff = width - diff
			}
		}
	}

	return buffer.String()
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 6))
}
