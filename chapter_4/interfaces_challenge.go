/*
	-- Interface Challenge --

- Write a struct called Capper that has a field to anthoer io.Writer and transforms everthing
	written to uppercase. Capper should implement io.Writer
*/
package main

import (
	"fmt"
	"io"
	"os"
)

// Capper is a capitalizer
type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	diff := byte('a' - 'A')

	out := make([]byte, len(p))

	for i, c := range p {
		if c >= 'a' && c <= 'z' {
			c -= diff
		}

		out[i] = c
	}

	return c.wtr.Write(out)
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello there")
}
