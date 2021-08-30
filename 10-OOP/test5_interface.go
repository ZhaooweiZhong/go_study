package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println(arg)
	value, ok := arg.(string)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("not string")
	}
}

type Book struct {
	auth string
}

func main() {
	num := 100
	myFunc(num)
	var book Book = Book{auth: "zzw"}
	myFunc(&book)
	str := "string"
	myFunc(str)
}
