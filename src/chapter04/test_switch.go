package main

func main() {
	grade := "B"
	switch grade {
	case "A":
		println("优秀")
	case "B":
		println("良好")
		fallthrough
	default:
		println("一般")
	}
	println("----------over----------")
}
