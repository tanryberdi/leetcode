package main

import "fmt"

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}

	numbers := make(map[string]int)
	numbers["I"] = 1
	numbers["V"] = 5
	numbers["X"] = 10
	numbers["L"] = 50
	numbers["C"] = 100
	numbers["D"] = 500
	numbers["M"] = 1000

	tail := 0
	result := 0
	for {
		if tail < len(s)-1 {
			if numbers[string(s[tail])] < numbers[string(s[tail+1])] {
				result += numbers[string(s[tail+1])] - numbers[string(s[tail])]
				tail += 2
			} else {
				result += numbers[string(s[tail])]
				tail++
			}
		} else {
			result += numbers[string(s[tail])]
			tail++
		}

		if tail >= len(s) {
			break
		}
	}

	return result
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
