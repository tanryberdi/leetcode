package main

import "fmt"

func calculateTax(brackets [][]int, income int) float64 {
	var result float64

	if income == 0 {
		return 0.0
	}

	if income < brackets[0][0] {
		result += float64(income*brackets[0][1]) / 100
		income = 0
	} else {
		result += float64(brackets[0][0]*brackets[0][1]) / 100
		income -= brackets[0][0]
	}

	for i := 1; i < len(brackets); i++ {
		if income < brackets[i][0]-brackets[i-1][0] {
			result += float64(income*brackets[i][1]) / 100
			income = 0
		} else {
			result += float64((brackets[i][0]-brackets[i-1][0])*brackets[i][1]) / 100
			income -= brackets[i][0] - brackets[i-1][0]
		}

		if income == 0 {
			break
		}
	}

	return result
}

func main() {
	brackets := [][]int{{1, 0}, {4, 25}, {5, 50}}
	income := 2
	fmt.Println(calculateTax(brackets, income))
}
