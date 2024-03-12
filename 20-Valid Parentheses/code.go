package main

import "fmt"

func main() {
	s := "({)({})})"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	var stack []string

	for _, v := range s {
		str := string(v)

		switch str {
		case "[":
			stack = append(stack, string(v))
		case "{":
			stack = append(stack, string(v))
		case "(":
			stack = append(stack, string(v))
		case "]":
			if len(stack) > 0 {
				if stack[len(stack)-1] == "[" {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		case "}":
			if len(stack) > 0 {
				if stack[len(stack)-1] == "{" {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		case ")":
			if len(stack) > 0 {
				if stack[len(stack)-1] == "(" {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		default:
			return false
		}
	}

	return len(stack) == 0
}
