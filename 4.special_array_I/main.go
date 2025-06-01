package main

import "fmt"

func main() {
	input := []int{2, 1, 4}
	a := isArraySpecial(input)

	fmt.Printf("the result is :%v", a)
}

func isArraySpecial(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i]%2 == nums[i+1]%2 {
			return false
		}
	}

	return true
}
