package main

import "fmt"

func main() {

	input := "abc7"
	a := clearDigits(input)

	fmt.Printf("the result is :%v", a)

}

func clearDigits(s string) string {
	stack := []rune{}

	if len(s) == 1 {
		return ""
	}

	for _, i := range s {
		if i >= '0' && i < '9' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, i)
		}
	}

	return string(stack)
}
