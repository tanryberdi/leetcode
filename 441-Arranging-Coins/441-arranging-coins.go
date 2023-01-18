package main

import "fmt"

func arrangeCoins(n int) int {
	k := 1
	sum := 2 * n
	presentSum := k * (k + 1)
	for presentSum <= sum {
		k++
		presentSum = k * (k + 1)
	}

	return k - 1
}

func main() {
	number := 3
	fmt.Println(arrangeCoins(number))
}
