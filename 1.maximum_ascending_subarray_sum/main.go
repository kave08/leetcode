package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 5, 10, 50}
	maxNum := maximumSum(nums)
	fmt.Println("sum max is :", maxNum)

}

func maximumSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	currentSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currentSum += nums[i]
		} else {
			maxSum = max(maxSum, currentSum)
			currentSum = nums[i]
		}
	}
	maxSum = max(maxSum, currentSum)

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
