package main

import "fmt"

func main() {
	var m1 map[string]string = make(map[string]string)
	for key, value := range m1 {
		fmt.Printf("%v=%v\n", key, value)
	}

}
