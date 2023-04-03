package main

func hammingWeight(num uint32) int {
	var count int
	for num != 0 {
		count++
		num &= num - 1
	}
	return count
}

func main() {
	println(hammingWeight(11111111111111111111111111111101))
}
