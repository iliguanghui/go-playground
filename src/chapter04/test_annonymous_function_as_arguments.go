// Go program to pass an anonymous
// function as an argument into
// other function
package main

import "fmt"

// Passing anonymous function
// as an argument
func GFG(i func(p, q string) string) {
	fmt.Printf("type of i is %T\n", i)
	fmt.Printf("content of i is %v\n", i)
	fmt.Printf("address of i is %p\n\n", &i)
	fmt.Println(i("Geeks", "for"))

}

func main() {
	value := func(p, q string) string {
		return p + q + "Geeks"
	}
	fmt.Printf("type of value is %T\n", value)
	fmt.Printf("content of value is %v\n", value)
	fmt.Printf("address of value is %p\n\n", &value)
	GFG(value)
}
