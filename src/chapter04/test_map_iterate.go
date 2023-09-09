package main

import "fmt"

func main() {
	var m1 = map[string]string{"name": "tom", "age": "20", "email": "tom@example.com"}
	for key := range m1 {
		fmt.Printf("%v=%v\n", key, m1[key])
	}
	println("--------------------")
	for key, value := range m1 {
		fmt.Printf("%v=%v\n", key, value)
	}
}
