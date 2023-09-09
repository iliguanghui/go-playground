package main

import "fmt"

func main() {
	var s1 []int
	fmt.Printf("%t\n", s1 == nil)
	fmt.Printf("%v\n", s1)
	s1 = append(s1, 100)
	fmt.Printf("%v\n", s1)
	fmt.Printf("%t\n", s1 == nil)
	s1[0] = 200
	fmt.Printf("%v\n", s1)
	fmt.Printf("%t\n", s1 == nil)
}
