package main

import (
	"fmt"
)

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	for {
		if len(s) == 0 {
			break
		}

		index := 0
		for i, c := range s {
			if string(c) == "(" || string(c) == "{" || string(c) == "[" {
				index = i
			}
		}

		if index == len(s)-1 {
			return false
		}

		switch string(s[index]) {
		case "(":
			if string(s[index+1]) == ")" {
				s = s[:index] + s[index+2:]
			} else {
				return false
			}
		case "[":
			if string(s[index+1]) == "]" {
				s = s[:index] + s[index+2:]
			} else {
				return false
			}
		default:
			if string(s[index+1]) == "}" {
				s = s[:index] + s[index+2:]
			} else {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(isValid("{[]}"))
	param := "{[]}"
	param = param[:1] + param[3:]
	fmt.Println(param)
}
