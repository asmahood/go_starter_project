// Parameter passing
package main

import "fmt"

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(i int) {
	i *= 2
}

func doublePtr(n *int) {
	*n *= 2
}

func main() {
	values := []int{1, 2, 3, 4}
	// when passing slices or maps, they are passed by reference
	doubleAt(values, 2)
	fmt.Println(values)

	val := 10
	// parameter is passed by value for other types
	double(val)
	fmt.Println(val)

	doublePtr(&val)
	fmt.Println(val)
}
