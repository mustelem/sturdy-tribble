package main

import "fmt"

func romanToInt(s string) int {
	var result int

	numerals := make(map[string]int)
	numerals["I"] = 1
	numerals["V"] = 5
	numerals["X"] = 10
	numerals["L"] = 50
	numerals["C"] = 100
	numerals["D"] = 500
	numerals["M"] = 1000

	for i := 0; i < len(s); i++ {
		thisChar := string(s[i])
		switch thisChar {
		case "I":
			if i < len(s)-1 {
				nextChar := string(s[i+1])
				if nextChar == "V" {
					result += 4
					i++
				} else if nextChar == "X" {
					result += 9
					i++
				} else {
					result += 1
				}
			} else {
				result += 1
			}
		case "X":
			if i < len(s)-1 {
				nextChar := string(s[i+1])
				if nextChar == "L" {
					result += 40
					i++
				} else if nextChar == "C" {
					result += 90
					i++
				} else {
					result += 10
				}
			} else {
				result += 10
			}
		case "C":
			if i < len(s)-1 {
				nextChar := string(s[i+1])
				if nextChar == "D" {
					result += 400
					i++
				} else if nextChar == "M" {
					result += 900
					i++
				} else {
					result += 100
				}
			} else {
				result += 100
			}

		default:
			result += numerals[string(s[i])]
		}
	}

	return result
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
