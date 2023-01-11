package main

import (
	"fmt"
	"strings"
)

func countSegments(s string) int {
	s = strings.Join(strings.Fields(s), " ")
	if s == "" {
		return 0
	}

	segments := strings.Split(s, " ")
	return len(segments)
}

func main() {
	st := ", , , ,        a, eaefa"
	fmt.Println(countSegments(st))
}
