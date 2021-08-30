package main

import "fmt"

func num1() {
	fmt.Println("num1")
}
func num2() {
	fmt.Println("num2")
}

func main() {
	defer num1()
	defer num2()
	fmt.Println("main")
}
