package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func isPalindrome(number int64) bool {
	if number < 10 {
		return true
	}

	if number%10 == 0 {
		return false
	}

	var newNumber int64
	temp := number
	newNumber = 0
	for temp > 0 {
		newNumber = newNumber*10 + (temp % 10)
		temp /= 10
	}

	if number == newNumber {
		return true
	}

	return false
}

// max_number is 999999999999999999
// sqrt of max_number = 1 000 000 000
func superpalindromesInRange(left string, right string) int {
	leftNumber, _ := strconv.Atoi(left)
	rightNumber, _ := strconv.Atoi(right)

	sqrtOfLeft := int(math.Round(math.Sqrt(float64(leftNumber))))
	sqrtOfRight := int(math.Round(math.Sqrt(float64(rightNumber))))

	counter := 0
	for i := sqrtOfLeft; i <= sqrtOfRight; i++ {
		if isPalindrome(int64(i)) {
			result := int64(i) * int64(i)
			if isPalindrome(result) {
				counter++
			}
		}
	}

	return counter
}

func main() {
	start := time.Now()
	fmt.Println(start)
	fmt.Println(superpalindromesInRange("38455498359", "999999999999999999"))
	fmt.Println(time.Since(start))
}
