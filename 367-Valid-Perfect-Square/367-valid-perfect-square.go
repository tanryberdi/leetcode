package main

import "fmt"

func isPerfectSquare(num int) bool {
	i, j := 1, num
	for i <= j {
		mid := i + (j-i)/2
		if mid*mid == num {
			return true
		}

		if mid*mid < num {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	return false
}

func main() {
	number := 25
	fmt.Println(isPerfectSquare(number))
}
