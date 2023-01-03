package main

import "fmt"

func addDigits(num int) int {
	for {
		sumOfDigits := 0
		for num > 0 {
			sumOfDigits += num % 10
			num /= 10
		}

		if sumOfDigits < 10 {
			return sumOfDigits
		}
		num = sumOfDigits
	}
}

func main() {
	number := 0
	fmt.Println(addDigits(number))
}
