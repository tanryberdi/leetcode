package main

import "fmt"

func getPermutation(n int, k int) string {
	var result string
	var nums []int
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	var factorial = make([]int, n+1)
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * i
	}
	k--
	for i := 1; i <= n; i++ {
		index := k / factorial[n-i]
		result += fmt.Sprint(nums[index])
		nums = append(nums[:index], nums[index+1:]...)
		k -= index * factorial[n-i]
	}
	return result
}

func main() {
	n := 3
	k := 3
	fmt.Println(getPermutation(n, k))

}
