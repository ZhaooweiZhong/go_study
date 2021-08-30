package main

import "fmt"

type Reader interface{
	Readbook()
}

type Writter interface{
	Writebook()
}

type Book struct{
	bookname string
}

func (this *Book)Readbook(){
	fmt.Println("Reading")
}
func (this *Book) Writebook(){
	fmt.Println("Writing")
}

func main(){
	b := &Book{"zzw's book"}
	var reader Reader
	reader = b
	reader.Readbook()
	var writer Writter
	writer = reader.(Writter)
	writer.Writebook()


}