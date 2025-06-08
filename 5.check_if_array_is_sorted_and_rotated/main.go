package main

import "fmt"

func main() {
	input := []int{3, 4, 5, 1, 2}
	a := check(input)

	fmt.Printf("the result is :%v", a)
}

func check(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	count := 0
	for i := range len(nums) - 1 {
		if nums[i] > nums[i+1] {
			count++
		}
	}

	if nums[len(nums)-1] > nums[0] {
		count++
	}

	return count <= 1
}
