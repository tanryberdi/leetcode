package main

import (
	"fmt"
	"strconv"
)

func convertToBase7(num int) string {
	var result string
	minus := false
	if num < 0 {
		minus = true
		num = -num
	}

	for {
		remainder := num % 7
		num /= 7
		remainderToString := strconv.Itoa(remainder)
		result = remainderToString + result
		if num == 0 {
			break
		}
	}

	if minus {
		return "-" + result
	}
	return result
}

func main() {
	number := -7
	fmt.Println(convertToBase7(number))
}
