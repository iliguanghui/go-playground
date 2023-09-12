// Go program to illustrate the
// concept of the call by value
package main

import "fmt"

// function which swap values
func swap(a, b int32) int32 {

	var o int32
	o = a
	a = b
	b = o
	fmt.Printf("%p,%p\n", &a, &b)

	return o
}

// Main function
func main() {
	var p int32 = 10
	var q int32 = 20
	fmt.Printf("p = %d and q = %d\n", p, q)
	fmt.Printf("%p,%p\n", &p, &q)

	// call by values
	swap(p, q)
	fmt.Printf("\np = %d and q = %d\n", p, q)
}
