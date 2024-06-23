package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}

func lengthOfLastWord(s string) int {
	var count int

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if count == 0 {
				count = 0
				continue
			} else {
				return count
			}
		} else {
			count++
		}
	}

	return count
}
