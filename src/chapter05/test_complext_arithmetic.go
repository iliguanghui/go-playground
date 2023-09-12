package main

import "fmt"

func main() {
	comp1 := complex(10, 11)
	comp2 := 13 + 33i
	fmt.Printf("%v\n", comp1)
	fmt.Printf("%v\n", comp2)
	realPart := real(comp1)
	imaginaryPart := imag(comp2)
	fmt.Printf("%v\n", realPart)
	fmt.Printf("%v\n", imaginaryPart)
}
