package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	res := []string{""}
	for i := 0; i < len(digits); i++ {
		temp := []string{}
		for j := 0; j < len(res); j++ {
			for k := 0; k < len(m[digits[i]]); k++ {
				temp = append(temp, res[j]+m[digits[i]][k])
			}
		}
		res = temp
	}
	return res
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
