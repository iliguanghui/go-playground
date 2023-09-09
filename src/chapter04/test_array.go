package main

import "fmt"

func main() {
	var a1 = [2]int{0, 1}
	var a2 = [...]string{0: "a", 2: "b", 4: "c"}
	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", a2)
	fmt.Printf("%v\n", a1)
	fmt.Printf("%v\n", a2)
}
