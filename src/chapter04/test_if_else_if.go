package main

func main() {
	if score := 85; score >= 60 && score < 70 {
		println("C")
	} else if score >= 70 && score < 80 {
		println("B")
	} else if score >= 80 && score < 90 {
		println("A")
	} else if score >= 90 {
		println("S")
	} else {
		println("D")
	}
}
