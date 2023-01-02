package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}

	if n <= 0 {
		return false
	}

	for {
		fmt.Println(n)
		if n == 1 {
			return true
		}

		if n%2 == 1 {
			return false
		}

		n /= 2
	}
	return true
}

func main() {
	number := -16
	fmt.Println(isPowerOfTwo(number))
}
