// If conditionals
package main

import (
	"fmt"
)

func main() {
	x := 3

	if x > 5 {
		fmt.Println("x is big")
	} else {
		fmt.Println("x is not big")
	}

	// Logical AND
	if x > 5 && x < 15 {
		fmt.Println("x is pretty cool")
	}

	// Logical OR
	if x < 4 || x > 6 {
		fmt.Println("x is not that cool")
	}

	// Assignment in if statement
	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}
}
