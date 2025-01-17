package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	var registry = make(map[string][]string)

	for _, word := range strs {
		runes := []rune(word)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		sortedWord := string(runes)

		v, ok := registry[sortedWord]
		if ok {
			registry[sortedWord] = append(v, word)
		} else {
			registry[sortedWord] = []string{word}
		}
	}

	var retVal = make([][]string, 0)

	for _, v := range registry {
		retVal = append(retVal, v)
	}

	return retVal
}

func main() {
	strs := []string{"ali", "lisa", "lia"}
	groupAnagrams(strs)
}
