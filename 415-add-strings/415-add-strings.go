package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	var result string

	if len(num1) == 0 {
		return num2
	}

	if len(num2) == 0 {
		return num1
	}

	if num1 == "0" {
		return num2
	}

	if num2 == "0" {
		return num1
	}

	i := len(num1) - 1
	j := len(num2) - 1

	sum := 0
	remainder := 0
	for {
		if i < 0 && j < 0 {
			break
		}

		//fmt.Println("i=", i)
		//fmt.Println("j=", j)

		if i >= 0 {
			lastDigitOfNum1, _ := strconv.Atoi(string(num1[i]))
			//fmt.Println("lastDigitOfNum1=", lastDigitOfNum1)
			sum += lastDigitOfNum1
			//fmt.Println("sum=", sum)
		}

		if j >= 0 {
			lastDigitOfNum2, _ := strconv.Atoi(string(num2[j]))
			//fmt.Println("lastDigitOfNum2=", lastDigitOfNum2)
			sum += lastDigitOfNum2
			//fmt.Println("sum=", sum)
		}

		sum += remainder
		remainder = 0
		if sum > 9 {
			sum %= 10
			remainder = 1
			//fmt.Println("sum=", sum)
			//fmt.Println("remainder=", remainder)
		}

		result = strconv.Itoa(sum) + result
		sum = 0
		//fmt.Println("result=", result)
		i--
		j--
	}

	if remainder > 0 {
		result = strconv.Itoa(remainder) + result
	}

	return result
}

func main() {
	num1 := "456"
	num2 := "77"
	fmt.Println(addStrings(num1, num2))
}
