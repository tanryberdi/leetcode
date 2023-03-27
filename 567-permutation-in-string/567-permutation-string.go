package main

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Map := make(map[rune]int)
	for _, r := range s1 {
		s1Map[r]++
	}
	for i := 0; i <= len(s2)-len(s1); i++ {
		s2Map := make(map[rune]int)
		for _, r := range s2[i : i+len(s1)] {
			s2Map[r]++
		}
		if isSame(s1Map, s2Map) {
			return true
		}
	}
	return false
}

func isSame(m1, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaoo"))
}
