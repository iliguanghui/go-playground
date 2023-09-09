package main

import "fmt"

func main() {
	var a = 100
	var b = 200
	if a > b {
		println(a)
	} else {
		println(b)
	}
	switch a {
	case 100:
		fmt.Printf("%d\n", 100)
		fallthrough
	case 200:
		fmt.Printf("%d\n", 200)
	default:
		fmt.Printf("%d\n", 1000)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	println("--------------------")
	x := [...]int{1, 2, 3, 4, 5}
	for _, i := range x {
		if i == 3 {
			break
		}
		fmt.Printf("%d\n", i)
	}
	println("-------end---------")
}
