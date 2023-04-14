package main

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		}
		if i+nums[i] > max {
			max = i + nums[i]
		}
		if max >= len(nums)-1 {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	println(canJump(nums))
}
