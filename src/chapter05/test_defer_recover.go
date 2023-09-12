package main

import "fmt"

func main() {
	demo(10)
	fmt.Println("程序继续执行")
}
func demo(index int) {
	var arr [10]int
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	arr[index] = 10
}
