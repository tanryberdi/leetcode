package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var result string
	var step int
	for i := 0; i < numRows; i++ {
		step = 2*numRows - 2
		for j := i; j < len(s); j += step {
			result += string(s[j])
			if i != 0 && i != numRows-1 && j+step-2*i < len(s) {
				result += string(s[j+step-2*i])
			}
		}
	}
	return result
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
