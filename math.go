package main

import "fmt"

func main() {
	fmt.Println(soma(112, 10))
}

func soma(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mult(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}
