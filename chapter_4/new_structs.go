// Methods in Go
package main

import (
	"fmt"
	"os"
)

// Trade is a trade in stocks
type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

// NewTrade will create a new trade and will validate the input
func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol cannot be empty")
	}

	if volume <= 0 {
		return nil, fmt.Errorf("volume must be >= 0 (was %d)", volume)
	}

	if price <= 0.0 {
		return nil, fmt.Errorf("price must be >= 0 (was %f)", price)
	}

	trade := &Trade{symbol, volume, price, buy}

	return trade, nil
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
	t, err := NewTrade("MSFT", 10, 99.89, true)

	if err != nil {
		fmt.Printf("error: cannot create trade - %s\n", err)
		os.Exit(1)
	}

	fmt.Println(t.Value())
}
