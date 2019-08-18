// For Loops
package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 3; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 3; i++ {
		if i > 1 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 3; i++ {
		if i < 1 {
			continue
		}
		fmt.Println(i)
	}

	// just a condition (like a while loop)
	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}

	// no condition ( while(true) )
	b := 0
	for {
		if b > 2 {
			break
		}
		fmt.Println(b)
		b++
	}

}
