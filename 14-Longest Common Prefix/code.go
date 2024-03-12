package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	strs := []string{"flower"}

	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ``
	}

	commonPre := strs[0]

	for _, word := range strs[1:] {
		if word == `` {
			return ``
		}

		commonPre = getPrefix(word, commonPre)

		if commonPre == `` {
			return ``
		}
	}

	return commonPre
}

func getPrefix(word string, pfx string) string {
	length := math.Min(float64(len(word)), float64(len(pfx)))
	lengthI := int(length)
	for i := lengthI; i > 0; i-- {
		if !strings.HasPrefix(word, pfx[:i]) {
			continue
		}

		return pfx[:i]
	}

	return ``
}
