// Go Slices

package main

import (
	"fmt"
)

func main() {
	// Same type
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	// length
	fmt.Println(len(loons))

	// 0 indexing
	fmt.Println(loons[1]) // daffy

	// slices
	fmt.Println(loons[1:])

	// -- LOOPING --

	// for
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	// single value range
	for i := range loons {
		fmt.Println(i)
	}

	// double value range
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	// double value range, ignore index using _
	for _, name := range loons {
		fmt.Println(name)
	}
}
