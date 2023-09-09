package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	var age int
	var email string
	fmt.Println("please input name, age, email: ")
	_, err := fmt.Scan(&name, &age, &email)
	if err != nil {
		os.Exit(1)
	}
	println("--------------------")
	fmt.Printf("%s\n", name)
	fmt.Printf("%d\n", age)
	fmt.Printf("%s\n", email)
}
