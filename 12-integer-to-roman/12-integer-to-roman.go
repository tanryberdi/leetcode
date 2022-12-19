package main

import "fmt"

func intToRoman(num int) string {
	if num < 1 && num > 3999 {
		return ""
	}

	var result string
	// for M's 1000
	for {
		if num >= 1000 {
			result += "M"
			num -= 1000
		} else {
			break
		}
	}

	// for CM's 900
	for {
		if num >= 900 {
			result += "CM"
			num -= 900
		} else {
			break
		}
	}

	// for D's 500
	for {
		if num >= 500 {
			result += "D"
			num -= 500
		} else {
			break
		}
	}

	// for CD's 400
	for {
		if num >= 400 {
			result += "CD"
			num -= 400
		} else {
			break
		}
	}

	// for C's 100
	for {
		if num >= 100 {
			result += "C"
			num -= 100
		} else {
			break
		}
	}

	// for XC's 90
	for {
		if num >= 90 {
			result += "XC"
			num -= 90
		} else {
			break
		}
	}

	// for L's 50
	for {
		if num >= 50 {
			result += "L"
			num -= 50
		} else {
			break
		}
	}

	// for XL's 40
	for {
		if num >= 40 {
			result += "XL"
			num -= 40
		} else {
			break
		}
	}

	// for X's 10
	for {
		if num >= 10 {
			result += "X"
			num -= 10
		} else {
			break
		}
	}

	// for IX's 9
	for {
		if num >= 9 {
			result += "IX"
			num -= 9
		} else {
			break
		}
	}

	// for V's 5
	for {
		if num >= 5 {
			result += "V"
			num -= 5
		} else {
			break
		}
	}

	// for IV's 4
	for {
		if num >= 4 {
			result += "IV"
			num -= 4
		} else {
			break
		}
	}

	// for I's 1
	for {
		if num >= 1 {
			result += "I"
			num -= 1
		} else {
			break
		}
	}

	return result
}

func main() {
	fmt.Println(intToRoman(3999))
}
