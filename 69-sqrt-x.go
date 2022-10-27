package main

import "fmt"

func mySqrt(x int) int {
	var result int
	for i := 0; i <= x; i++ {
		if i*i > x {
			break
		}
		result = i
	}
	return result
}

func main() {
	x := 4
	fmt.Println(mySqrt(x))
}
