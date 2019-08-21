/*
	FizzBuzz Challenge
loop through the numbers 1-20
- if the number is divsible by 3, print "fizz"
- if the number is divisble by 5, print "buzz"
- if the number is divisible by 3 and 5, print "fizzbuzz"

if not, print the number
*/

package main

import (
	"fmt"
)

func main() {
	// for every number between 1 and 20
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			// number is divisble by 3 and 5
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			// number is divisble by 3
			fmt.Println("fizz")
		} else if i%5 == 0 {
			// number is divisble by 5
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
