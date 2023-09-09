package main

import (
	"fmt"
)

func main() {
	name = "zhangsan"
	fmt.Println("name = " + name)
	age := 18
	fmt.Println("age = ", age)
	fmt.Printf("%T, %T\n", name, age)
	println("----------------")
	a := 10
	b := 20
	c, d := swap(a, b)
	fmt.Println(a, b, c, d)
	fmt.Printf("变量a的地址是:%x\n", &a)
	fmt.Printf("变量b的地址是:%x\n", &b)
	fmt.Printf("变量c的地址是:%x\n", &c)
	fmt.Printf("变量d的地址是:%x\n", &d)
	fmt.Println("--------------------")
	swap2(&a, &b)
	fmt.Println(a, b, c, d)
}

var name string

func swap2(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func swap(x, y int) (int, int) {
	fmt.Printf("变量x的地址是:%x\n", &x)
	fmt.Printf("变量y的地址是:%x\n", &y)
	var temp int
	temp = x
	x = y
	y = temp
	return x, y
}
