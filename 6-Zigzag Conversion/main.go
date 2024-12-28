package main

import (
	"bytes"
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	var buffer bytes.Buffer
	letters := strings.Split(s, "")

	width := numRows + max(0, numRows-2)
	diff := width

	for j := 0; j < numRows; j++ {
		i := j

		diff = width - 2*j
		if diff == 0 {
			diff = width - diff
		}

		for i < len(letters) {
			buffer.WriteString(letters[i])
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
