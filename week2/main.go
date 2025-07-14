package main

import "fmt"

func swap(a, b *string) {
	*a, *b = *b, *a
}

func max(s []int) int {
	if len(s) == 0 {
		return 0
	}

	v := s[0]

	if len(s) == 1 {
		return v
	}

	for i := 1; i < len(s); i++ {
		if v < s[i] {
			v = s[i]
		}
	}

	return v
}

func findEvenNumber(v []int) ([]int, error) {
	r := []int{}

	for i := 0; i < len(v); i++ {
		if v[i]%2 == 0 {
			r = append(r, v[i])
		}
	}

	if len(r) == 0 {
		return []int{}, fmt.Errorf("no even number found in slice")
	}

	return r, nil
}

func findOddNumber(v []int) ([]int, error) {
	r := []int{}
	for i := 0; i < len(v); i++ {
		if v[i]%2 != 0 {
			r = append(r, v[i])
		}
	}

	if len(r) == 0 {
		return []int{}, fmt.Errorf("no odd number found in slice")
	}

	return r, nil
}

func main() {
	x := "a"
	y := "b"
	fmt.Println("Original Values = ", x, y)
	swap(&x, &y)
	fmt.Println("Swapped Values = ", x, y)

	fmt.Println("")

	values := []int{1, 2, 3, 4, 5}
	maxValue := max(values)
	fmt.Println("Max Value in", values, "=", maxValue)

	fmt.Println("")

	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n1, _ := findEvenNumber(s1)
	fmt.Println("Even Numbers in", s1, "=", n1)
	fmt.Println("")

	o1 := []int{1, 3, 5, 7, 9}
	_, err := findEvenNumber(o1)
	if err != nil {
		fmt.Println(err, o1)
	}
	fmt.Println("")

	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n2, _ := findOddNumber(s2)
	fmt.Println("Odd Numbers in", s2, "=", n2)
	fmt.Println("")

	o2 := []int{2, 4, 6, 8, 10}
	_, err2 := findOddNumber(o2)
	if err2 != nil {
		fmt.Println(err, o2)
	}
}
