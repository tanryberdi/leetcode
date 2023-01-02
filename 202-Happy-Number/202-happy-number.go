package main

import "fmt"

func isHappy(n int) bool {
	if n < 0 {
		return false
	}

	numbers := make(map[int]int)
	for {
		numbers[n]++
		if n == 1 {
			break
		}

		newNumber := 0
		for n > 0 {
			remainder := n % 10
			newNumber += remainder * remainder
			n /= 10
		}
		n = newNumber
		numbers[n]++
		if numbers[n] > 1 {
			return false
		}
		fmt.Println(newNumber)
	}

	return true
}

func main() {
	number := 7
	fmt.Println(isHappy(number))
}
