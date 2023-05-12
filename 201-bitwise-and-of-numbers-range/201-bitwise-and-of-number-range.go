package main

import "fmt"

func rangeBitwiseAnd(left int, right int) int {
	shift := 0
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}
	return left << shift
}

func main() {
	left := 1
	right := 2147483647
	fmt.Println(rangeBitwiseAnd(left, right))
}
