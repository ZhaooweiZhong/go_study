package main

import "fmt"

func foll(a string, b int) int {
	fmt.Println(a)
	fmt.Println(b)
	c := 100
	return c
}

func foll2(a string, b int) (c int, d string) {
	fmt.Println(a)
	fmt.Println(b)
	c = 100
	d = "abbb"
	return
}

func main() {
	c, d := foll2("aaa", 2)
	fmt.Println(c, d)
}
