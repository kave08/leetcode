package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4, 5, 6}
	a := applyOperations(input)
	fmt.Printf("result is %v", a)

}

func applyOperations(nums []int) []int {
	n := len(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}

	}

	pos := 0
	for i := range n {
		if nums[i] != 0 {
			nums[pos] = nums[i]
			pos++
		}
	}

	for i := pos; i < n; i++ {
		nums[i] = 0
	}

	return nums

}
