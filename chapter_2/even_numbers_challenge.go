/*
--	Even-Ended Numbers Challenge --

A even-ended number is a number with the same first and last digits
	ex: 1, 11, 121

How many even-ended numbers result from multiplying two four-digit numbers?

*/

package main

import (
	"fmt"
)

func main() {
	// Sprintf example
	n := 42
	s := fmt.Sprintf("%d", n)

	fmt.Printf("s = %q (type %T)\n", s, s)

	// Problem starts here

	// counter variable
	count := 0

	// loop through first 4 digit numbers
	for x := 1000; x <= 9999; x++ {
		// loop through second 4 digit numbers, dont double count
		for y := x; y <= 9999; y++ {
			prod := x * y

			// convert prod to a string
			s := fmt.Sprintf("%d", prod)

			// if its even ended, increment count
			if s[0] == s[len(s)-1] {
				count++
			}
		}
	}

	// print count
	fmt.Println(count)
}
