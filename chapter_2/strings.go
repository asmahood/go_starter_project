// Go strings

package main

import (
	"fmt"
)

func main() {
	book := "the colour of magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	// uint8 is a byte

	// Strings are immutable
	// book[0] = 116

	// Slice (start, end), 0 based, 1/2 empty range
	fmt.Println(book[4:11])

	// Slice (no end)
	fmt.Println(book[4:])

	// Slice (no start)
	fmt.Println(book[:4])

	// Use + to concatentate strings
	fmt.Println("f" + book[1:])

	// Unicode
	fmt.Println("It was ½ price!")

	// multi line
	poem := `
	The road goes ever on
	Further down the road
	`

	fmt.Println(poem)
}
