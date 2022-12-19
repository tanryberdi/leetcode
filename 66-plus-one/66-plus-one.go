package main

import "fmt"

func plusOne(digits []int) []int {
	digits[len(digits)-1]++
	if digits[len(digits)-1] == 10 {
		digits[len(digits)-1] = 0
		inMemory := 1

		for i := len(digits) - 2; i >= 0; i-- {
			digits[i] += inMemory
			if digits[i] == 10 {
				digits[i] = 0
				inMemory = 1
			} else {
				inMemory = 0
			}

			if inMemory == 0 {
				break
			}
		}

		if inMemory == 1 {
			digits = append(digits, 0)
			copy(digits[1:], digits)
			digits[0] = 1
		}
	}

	return digits
}

func main() {
	digits := []int{8, 9, 9, 9}
	fmt.Println(plusOne(digits))
}
