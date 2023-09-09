package main

import "fmt"

func main() {
	var number1 = 100
	var number2 = 200
	var number3 = 300
	if number1 > number2 {
		if number1 > number3 {
			fmt.Printf("largest number is %d\n", number1)
		} else {
			fmt.Printf("largest number is %d\n", number3)
		}
	} else {
		if number2 > number3 {
			fmt.Printf("largest number is %d\n", number2)
		} else {
			fmt.Printf("largest number is %d\n", number3)
		}
	}
}
