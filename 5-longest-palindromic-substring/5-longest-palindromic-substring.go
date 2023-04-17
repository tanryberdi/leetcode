package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	maxLen := 1
	begin := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if j-i+1 > maxLen && validPalindrome(s, i, j) {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

func validPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "cbbd"
	fmt.Println(longestPalindrome(s))
}
