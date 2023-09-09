package main

import "fmt"

func main() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", a)
	for i, v := range a {
		fmt.Printf("i: %d, v = %d\n", i, v)
	}

	println("--------------------")
	var b = []string{"alpha", "beta", "gamma"}
	for _, s := range b {
		println(s)
	}
	println("--------------------")
	var m = make(map[string]string)
	m["name"] = "tom"
	m["age"] = "20"
	m["email"] = "test@example.com"
	for key, value := range m {
		fmt.Printf("%v=%v\n", key, value)
	}
}
