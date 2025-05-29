package main

import "fmt"

func main() {

	s1 := "bank"
	s2 := "kanb"
	a := areAlmostEqual(s1, s2)
	fmt.Println("the strings are equal :", a)

}

func areAlmostEqual(s1, s2 string) bool {

	if s1 == s2 {
		return true
	}

	var diffIndices []int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffIndices = append(diffIndices, i)
		}
	}

	if len(diffIndices) != 2 {
		return false
	}

	i, j := diffIndices[0], diffIndices[1]

	return s1[i] == s2[j] && s1[j] == s2[i]
}
