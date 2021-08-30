package main

import "fmt"

func example(p *int) {
	fmt.Println(*p)
}
func main() {
	a := 10
	example(&a)
}
