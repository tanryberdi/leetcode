package main

import (
	"fmt"
	"strings"
)

func findTheDifference(s string, t string) byte {
	for _, char := range t {
		index := strings.Index(s, string(char))
		if index == -1 {
			return byte(char)
		}
		s = s[:index] + s[index+1:]
	}
	return 0
}

func main() {
	st1 := "a"
	st2 := "aa"
	fmt.Println(findTheDifference(st1, st2))
}
