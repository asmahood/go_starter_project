// Go Garbage Collector
package main

import "fmt"

func cleanup(name string) {
	fmt.Printf("cleaning up %s\n", name)
}

func worker() {
	defer cleanup("A")

	fmt.Println("worker")
}

func main() {
	worker()
}
