package main

import "fmt"

type Website struct {
	name string
	url  string
}

func main() {
	var site = Website{
		name: "google",
		url:  "www.google.com"}
	fmt.Printf("site = %#v\n", site)
	fmt.Printf("site = %T\n", site)
	var a int = 100
	fmt.Printf("%T\n", a)
	i := 8
	fmt.Printf("i: %v\n", i)
	fmt.Printf("i: %b\n", i)
	fmt.Printf("i: %o\n", i)
	i = 0x61
	fmt.Printf("i: %c\n", i)
	fmt.Printf("%p\n", &i)
}
