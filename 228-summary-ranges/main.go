package main

import "fmt"

func summaryRanges(nums []int) []string {
	var res []string
	for i := 0; i < len(nums); i++ {
		j := i
		for j+1 < len(nums) && nums[j+1] == nums[j]+1 {
			j++
		}
		if j > i {
			res = append(res, fmt.Sprintf("%d->%d", nums[i], nums[j]))
		} else {
			res = append(res, fmt.Sprintf("%d", nums[i]))
		}
		i = j
	}
	return res
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	result := summaryRanges(nums)
	fmt.Println(result)
}
