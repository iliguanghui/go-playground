// Go program to illustrate the
// use of function
package main

import "fmt"

// area() is used to find the
// area of the rectangle
// area() function two parameters,
// i.e, length and width
func getArea(length, width int) int {

	area := length * width
	return area
}

// Main function
func main() {

	// Display the area of the rectangle
	// with method calling
	fmt.Printf("Area of rectangle is : %d", getArea(12, 10))
}
