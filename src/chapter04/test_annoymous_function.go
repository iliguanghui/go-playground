// Go program to illustrate how
// to create an anonymous function
package main

import "fmt"

func main() {

	// Anonymous function
	func() {
		fmt.Println("Welcome! to GeeksforGeeks")
	}()
	var name = "jack"
	fmt.Printf("address of name is %p\n", &name)

	// Assigning an anonymous
	// function to a variable
	value := func(name string) int {
		fmt.Println(name + ", welcome to GeeksforGeeks")
		return 10
	}
	fmt.Printf("type of value is %T\n", value)
	fmt.Printf("address of value is %p\n", &value)
	fmt.Printf("%v\n", value)
	value("jenkins")
}
