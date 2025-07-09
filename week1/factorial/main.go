package main

import "fmt"

func main() {
	var n uint32 = 5
	var factorial uint32 = 1

	if n == 0 {
		factorial = 1
	} else {
		factorial = n
		for i := n - 1; i > 0; i-- {
			factorial *= i
		}
	}

	fmt.Println("Factorial of", n, "is", factorial)
}
