package main

import (
	"bytes"
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	var activeWidth int
	letters := strings.Split(s, "")

	L := len(letters)

	dist := max(numRows-2, 0)
	quot1 := L / (numRows + dist)
	rem1 := L % (numRows + dist)
	quot2 := rem1 / numRows
	rem2 := rem1 % numRows

	w := quot1*(dist+1) + quot2 + rem2

	var table [][]string = make([][]string, numRows)

	for i := range table {
		table[i] = make([]string, w) // TODO Not sure about the width, yet.
	}

	i := 0
	j := 0
	letterPtr := 0

	down := true
	for {
		table[i][j] = letters[letterPtr]
		letterPtr++
		if letterPtr > len(letters)-1 {
			activeWidth = j
			break
		}

		if down {
			i++
			if i > numRows-1 {
				if i-2 < 0 {
					i = 0
				} else {
					i = i - 2
					down = false
				}
				j++
			}
		} else {
			i--
			j++
			if i < 0 {
				down = true
				j--
				i = i + 2
			}
		}
	}

	var buffer bytes.Buffer
	for i := 0; i < numRows; i++ {
		for j := 0; j <= activeWidth; j++ {
			if table[i][j] != "" {
				buffer.WriteString(table[i][j])
			}
		}
	}

	return buffer.String()
}

func main() {
	for i := 3; i < 9; i++ {
		fmt.Println(convert("PAYPALISHIRING", i))
	}
}
