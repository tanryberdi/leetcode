package main

func isUgly(n int) bool {
	if n == 0 {
		return false
	}
	for _, i := range []int{2, 3, 5} {
		for n%i == 0 {
			n /= i
		}
	}
	return n == 1
}

func main() {
	println(isUgly(6))
	println(isUgly(8))
	println(isUgly(14))
	println(isUgly(1))
	println(isUgly(0))
	println(isUgly(-6))
}
