/*
	-- Struct Challenge --

- Define a square struct which has two fields: center of type Point
	and length of type int. Add two methods:
		- Move(dx, dy)
		- Area() int
- Also write
	- NewSquare(x, y, length) (*Square, error)
*/
package main

import (
	"fmt"
	"log"
)

// Point is a 2D point
type Point struct {
	X int
	Y int
}

// Move moves the point by dx and dy
func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

// Square is a 2D square
type Square struct {
	Center Point
	Length int
}

// Move moves the sqaure by dx and dy
func (s *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}

// Area returns the area of the square
func (s *Square) Area() int {
	return s.Length * s.Length
}

// NewSquare creates a new square
func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be > 0 (was %d)", length)
	}

	s := &Square{
		Center: Point{x, y},
		Length: length,
	}

	return s, nil
}

func main() {
	s, err := NewSquare(1, 1, 5)

	if err != nil {
		log.Fatalf("ERROR: can't create square - %s", err)
	}

	s.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())
}
