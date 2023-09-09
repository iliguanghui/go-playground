package main

func main() {
	var gender = "男"
	var age int = 20
	if gender == "男" {
		if age >= 18 {
			println("成年男性")
		} else {
			println("未成年男性")
		}
	} else {
		if age >= 18 {
			println("成年女性")
		} else {
			println("未成年女性")
		}
	}
}
