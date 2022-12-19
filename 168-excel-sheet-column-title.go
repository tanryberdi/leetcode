package main

import "fmt"

func convertToTitle(columnNumber int) string {
	if columnNumber == 0 {
		return ""
	}
	bytes := []byte{}
	var remainder int
	for columnNumber > 0 {
		remainder = (columnNumber - 1) % 26
		bytes = append(bytes, 'A'+byte(remainder))
		columnNumber = (columnNumber - 1) / 26
	}
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}

func main() {
	columnNumber := 2147483647
	fmt.Println(convertToTitle(columnNumber))
}
