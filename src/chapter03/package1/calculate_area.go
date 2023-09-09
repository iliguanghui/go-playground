package main

import (
	"fmt"
)

func main() {
	const length int = 10
	const width int = 5
	area := length * width
	const a, b, c = 1, false, "string"

	fmt.Printf("area = %d\n", area)
	area = length * width * 2

	fmt.Printf("2area = %d\n", area)
	fmt.Println(a, b, c)

	fmt.Println(Unknown, Female, Male)
}

const (
	Unknown = 0 + iota
	Female
	Male
)
