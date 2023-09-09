package main

import "fmt"

func main() {
	var a1 [2]int
	fmt.Printf("%v\n", a1)
	a1[0] = 100
	a1[1] = 200
	fmt.Printf("%v\n", a1)

	var s1 []int
	var s2 []string
	fmt.Printf("%T\n", s1)
	fmt.Printf("%T\n", s2)
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
	fmt.Printf("%v\n", s1 == nil)
	fmt.Printf("%v\n", s2 == nil)
}
