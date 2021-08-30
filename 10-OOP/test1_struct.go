package main

import "fmt"

type Book struct {
	title string
	auth  string
}

func changebook(bookin Book) {
	bookin.title = "title have been changed"
}

func changebook2(bookin *Book) {
	bookin.title = "title have been changed"
}

func main() {
	var book1 Book
	book1.title = "Hello World"
	book1.auth = "zzw"

	fmt.Println(book1)
	changebook(book1)
	fmt.Println(book1)
	changebook2(&book1)
	fmt.Println(book1)
}
