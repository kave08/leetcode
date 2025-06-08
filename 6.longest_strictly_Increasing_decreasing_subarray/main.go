package main

import "fmt"

func main() {

	input := []int{1, 4, 3, 3, 2}
	a := longestMonotonicSubarray(input)

	fmt.Printf("the result is :%v", a)

}

func longestMonotonicSubarray(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	incLen := 1
	decLen := 1
	maxLen := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			incLen++
			decLen = 1
		} else if nums[i] < nums[i-1] {
			decLen++
			incLen = 1
		} else {
			decLen = 1
			incLen = 1
		}
		maxLen = max(maxLen, max(incLen, decLen))
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
