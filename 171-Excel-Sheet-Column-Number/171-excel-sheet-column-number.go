package main

import (
	"fmt"
)

func titleToNumber(columnTitle string) int {
	result := 0

	for i := 0; i < len(columnTitle); i++ {
		last := int(columnTitle[i] - 'A' + 1)
		result = result*26 + last
	}

	return result
}

func main() {
	columnTitle := "ZY"
	fmt.Println(titleToNumber(columnTitle))
}
