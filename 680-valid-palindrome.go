package main

import "fmt"

func validPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	tail := 0
	head := len(s) - 1

	// nod is a number of deleted characters from string
	nod := 0

	result := true
	for {
		if head < tail {
			break
		}

		// if they are same characters
		if s[tail] == s[head] {
			tail++
			head--
			continue
		} else {
			if nod > 0 {
				result = false
				head = tail - 1
			}

			if head < tail {
				break
			}

			if s[tail] == s[head-1] {
				tail++
				head -= 2
				nod++
				continue
			}

			if s[tail+1] == s[head] {
				tail += 2
				head--
				nod++
				continue
			}

			result = false
			head = tail - 1
		}
	}

	if result {
		return true
	}

	tail = 0
	head = len(s) - 1

	// nod is a number of deleted characters from string
	nod = 0

	for {
		if head < tail {
			break
		}

		// if they are same characters
		if s[tail] == s[head] {
			tail++
			head--
			continue
		} else {
			if nod > 0 {
				return false
			}

			if s[tail+1] == s[head] {
				tail += 2
				head--
				nod++
				continue
			}

			if s[tail] == s[head-1] {
				tail++
				head -= 2
				nod++
				continue
			}

			return false
		}
	}

	return true
}

func main() {
	s := "aeacdeaeaaaaaaeaedcae"
	fmt.Println(validPalindrome(s))
}
