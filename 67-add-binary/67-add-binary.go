package main

import "fmt"

func addBinary(a string, b string) string {
	var result string
	i := len(a) - 1
	j := len(b) - 1

	result = ""
	var sum, remainder uint8
	for {
		sum = 0
		if i < 0 && j < 0 {
			break
		}

		if i >= 0 {
			sum += a[i] - 48
			i--
		}

		if j >= 0 {
			sum += b[j] - 48
			j--
		}

		sum += remainder
		remainder = sum / 2
		sum %= 2

		result = string(48+sum) + result
	}

	if remainder == 1 {
		result = "1" + result
		sum = 0
	}

	return result
}

func main() {
	fmt.Println(addBinary("11", "1"))
}
