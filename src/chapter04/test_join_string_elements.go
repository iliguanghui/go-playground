// Go program to illustrate the
// concept of variadic function
package main

import "fmt"
import "strings"

// Variadic function to join strings
func joinstr(elements ...string) string {
	fmt.Printf("type of elements is %T\n", elements)
	fmt.Printf("address of elements is %p\n", &elements)
	if elements == nil {
		fmt.Println("elements is nil")
	}
	return strings.Join(elements, "-")
}

func main() {
	var a = 100
	fmt.Printf("address of variable a is %p\n", &a)

	fmt.Println("--------------------")
	// zero argument
	fmt.Println(joinstr())

	// multiple arguments
	fmt.Println("--------------------")
	fmt.Println(joinstr("GEEK", "GFG"))
	fmt.Println("--------------------")
	fmt.Println(joinstr("Geeks", "for", "Geeks"))
	fmt.Println("--------------------")
	fmt.Println(joinstr("G", "E", "E", "k", "S"))

	// pass a slice in variadic function
	elements := []string{"geeks", "FOR", "geeks"}
	// unpack slice
	fmt.Println(joinstr(elements...))
}
