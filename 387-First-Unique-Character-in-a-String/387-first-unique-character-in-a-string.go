package main

import "fmt"

func firstUniqChar(s string) int {
	letters := make(map[int32]int)
	for _, char := range s {
		letters[char]++
	}

	for index, char := range s {
		if letters[char] == 1 {
			return index
		}
	}

	return -1
}

func main() {
	st := "aabb"
	fmt.Println(firstUniqChar(st))
}
