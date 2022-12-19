package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	newString := ""
	for _, ch := range s {
		if unicode.IsDigit(ch) || unicode.IsLetter(ch) {
			newString += strings.ToLower(string(ch))
		}
	}

	fmt.Println(newString)

	if len(newString) == 0 {
		return true
	}

	tail := 0
	head := len(newString) - 1
	for {
		if head < tail {
			break
		}

		if newString[tail] != newString[head] {
			return false
		}

		tail++
		head--
	}

	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	//s := ""
	fmt.Println(isPalindrome(s))
}
