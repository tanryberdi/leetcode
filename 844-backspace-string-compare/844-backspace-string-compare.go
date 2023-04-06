package main

import "fmt"

func backspaceCompare(S string, T string) bool {
	s := []rune(S)
	t := []rune(T)
	s = backspace(s)
	t = backspace(t)
	return string(s) == string(t)
}

func backspace(s []rune) []rune {
	var stack []rune
	for _, v := range s {
		if v == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, v)
		}
	}
	return stack
}

func main() {
	s := "ab#c"
	t := "ad#c"
	fmt.Println(backspaceCompare(s, t))
}
