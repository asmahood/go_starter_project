// Methods in Go
package main

import "fmt"

// Trade is a trade in stocks
type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

// Value returns the trade value
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price

	if t.Buy {
		value = -value
	}

	return value
}

func main() {
	t := Trade{"MSFT", 10, 98.99, true}
	fmt.Println(t.Value())
}
