package main

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	var res []string
	i := 0
	for {
		j := i
		for j+1 < len(nums) && nums[j+1] == nums[j]+1 {
			j++
		}
		if j > i {
			res = append(res, fmt.Sprintf("%d->%d", nums[i], nums[j]))
		} else {
			res = append(res, fmt.Sprintf("%d", nums[i]))
		}

		// We need to move i to j+1
		i = j + 1

		if i == len(nums) {
			break
		}
	}
	return res
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	result := summaryRanges(nums)
	fmt.Println(result)
}
