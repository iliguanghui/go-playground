package main

import (
	"fmt"
)

func main() {
	var a = 100
	var b = 200
	fmt.Printf("%d\n", a+b)
	a++
	print(a)
	a = 4
	b = 8
	fmt.Printf("a = %b\n", a) // 0100
	fmt.Printf("b = %b\n", b) // 1000
	fmt.Printf("%b\n", a&b)
	fmt.Printf("%b\n", a|b)
	fmt.Printf("%b\n", a^b)
	a = -256
	fmt.Printf("%b\n", a>>2)
	fmt.Printf("%d\n", getANumber())
}

func getANumber() (number int) {
	number = 131
	return
}
