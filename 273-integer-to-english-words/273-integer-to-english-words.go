package main

import (
	"fmt"
	"strings"
)

func numberToWords(num int) string {
	numbers := make(map[int]string)

	if num < 0 {
		return ""
	}

	numbers[0] = "Zero"
	if num == 0 {
		return numbers[num]
	}
	var result string

	numbers[1] = "One"
	numbers[2] = "Two"
	numbers[3] = "Three"
	numbers[4] = "Four"
	numbers[5] = "Five"
	numbers[6] = "Six"
	numbers[7] = "Seven"
	numbers[8] = "Eight"
	numbers[9] = "Nine"
	numbers[10] = "Ten"
	numbers[11] = "Eleven"
	numbers[12] = "Twelve"
	numbers[13] = "Thirteen"
	numbers[14] = "Fourteen"
	numbers[15] = "Fifteen"
	numbers[16] = "Sixteen"
	numbers[17] = "Seventeen"
	numbers[18] = "Eighteen"
	numbers[19] = "Nineteen"
	numbers[20] = "Twenty"
	numbers[30] = "Thirty"
	numbers[40] = "Forty"
	numbers[50] = "Fifty"
	numbers[60] = "Sixty"
	numbers[70] = "Seventy"
	numbers[80] = "Eighty"
	numbers[90] = "Ninety"
	numbers[100] = "Hundred"
	numbers[1000] = "Thousand"
	numbers[1000000] = "Million"
	numbers[1000000000] = "Billion"

	for {
		if num >= 1000000000 {
			result += numbers[num/1000000000] + " Billion "
			num %= 1000000000
		}

		if num >= 1000000 {
			div := num / 1000000
			for div > 0 {
				if div >= 100 {
					result += numbers[div/100] + " Hundred "
					div %= 100
				}

				if div >= 21 {
					result += numbers[(div/10)*10] + " "
					div %= 10
				}

				if div > 0 && div < 21 {
					result += numbers[div] + ""
					div = 0
				}

				if div > 0 {
					result += numbers[div] + " "
					div = 0
				}

				result += " Million "
			}

			num %= 1000000
		}

		if num >= 1000 {
			div := num / 1000
			for div > 0 {
				if div >= 100 {
					result += numbers[div/100] + " Hundred "
					div %= 100
				}

				if div >= 21 {
					result += numbers[(div/10)*10] + " "
					div %= 10
				}

				if div > 0 && div < 21 {
					result += numbers[div] + ""
					div = 0
				}

				if div > 0 {
					result += numbers[div] + " "
					div = 0
				}

				result += " Thousand "
			}

			num %= 1000
		}

		if num > 0 {
			for num > 0 {
				if num >= 100 {
					result += numbers[num/100] + " Hundred "
					num %= 100
				}

				if num >= 21 {
					result += numbers[(num/10)*10] + " "
					num %= 10
				}

				if num > 0 && num < 21 {
					result += numbers[num] + ""
					num = 0
				}

				if num > 0 {
					result += numbers[num] + " "
					num = 0
				}
			}
		}

		break

	}

	result = strings.TrimSpace(strings.Replace(result, "  ", " ", -1))
	return result
}

func main() {
	//fmt.Println(numberToWords(2147483647))
	fmt.Println(numberToWords(50868))
}
