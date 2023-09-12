// Go program to illustrate
// use of anonymous function
package main

import "fmt"

// GFG2 Returning anonymous function
func GFG2() func(i, j string) string {
	myf := func(i, j string) string {
		return i + j + "GeeksforGeeks"
	}
	return myf
}

func main() {
	value := GFG2()
	fmt.Printf("type of value is %T\n", value)
	fmt.Println(value("Welcome ", "to "))
}
