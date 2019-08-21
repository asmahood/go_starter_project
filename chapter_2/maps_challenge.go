/*
	-- Word Count Challenge

Count the number of words in a text
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind`

	wc := map[string]int{}

	words := strings.Fields(text)

	for _, word := range words {
		/* SOLUTION
		wc[strings.ToLower(word)]++
		*/
		word = strings.ToLower(word)

		value, ok := wc[word]

		if !ok {
			wc[word] = 1
		} else {
			wc[word] = value + 1
		}
	}

	fmt.Println(wc)
}
