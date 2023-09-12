// Go program to illustrate the
// concept of the call by reference
package main

import "fmt"

// function which swap values
func swap2(a, b *int) int {
	var o int
	o = *a
	*a = *b
	*b = o
	fmt.Printf("%p,%p\n", a, b)

	return o
}

// Main function
func main() {

	var p int = 10
	var q int = 20
	fmt.Printf("p = %d and q = %d\n", p, q)
	fmt.Printf("%p,%p\n", &p, &q)

	// call by reference
	swap2(&p, &q)
	fmt.Printf("\np = %d and q = %d\n", p, q)
}
