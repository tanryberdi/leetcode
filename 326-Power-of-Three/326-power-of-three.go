package main

import "fmt"

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}

	if n <= 0 {
		return false
	}

	for {
		if n == 1 {
			return true
		}

		if n%3 != 0 {
			return false
		}

		n /= 3
	}
}

func main() {
	number := 27
	fmt.Println(isPowerOfThree(number))
}
