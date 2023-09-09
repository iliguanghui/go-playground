package main

import (
	"fmt"
	"os"
)

func main() {
	var c string
	fmt.Println("请输入一个字符:")
	_, err := fmt.Scanf("%s", &c)
	if err != nil {
		os.Exit(1)
	}
	if c == "M" {
		println("Monday")
	} else if c == "T" {
		fmt.Println("请输入第二个字符")
		_, err = fmt.Scanf("%s", &c)
		if err != nil {
			os.Exit(1)
		}
		if c == "u" {
			println("Tuesday")
		} else {
			println("Thursday")
		}
	} else if c == "W" {
		println("Wednesday")
	} else if c == "F" {
		println("Friday")
	} else if c == "S" {
		fmt.Println("请输入第二个字符")
		_, err = fmt.Scanf("%s", &c)
		if err != nil {
			os.Exit(1)
		}
		if c == "a" {
			println("Saturday")
		} else {
			println("Sunday")
		}
	}
}
