package main

import "fmt"

func main() {
	s := []int{123, 213, 444, 123}

	n := len(s)
	maxValue := s[0]

	for i := 0; i < n-1; i++ {
		if maxValue < s[i] {
			maxValue = s[i]
		}
	}

	fmt.Println(maxValue, "is the max value of", s)
}
