package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)

	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 32 {
			length++
		} else {
			break
		}
	}

	return length
}

func main() {
	s := "luffy is still joyboy"
	fmt.Println(lengthOfLastWord(s))
}
