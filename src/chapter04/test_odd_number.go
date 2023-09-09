package main

import (
	"fmt"
	"os"
)

func main() {
	var number int
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		os.Exit(1)
	}
	if number%2 != 0 {
		fmt.Printf("%d is an odd number", number)
	} else {
		fmt.Printf("%d is an even number", number)
	}
}
