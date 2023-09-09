package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println("Hello Go!")
	os.Exit(0)
}
