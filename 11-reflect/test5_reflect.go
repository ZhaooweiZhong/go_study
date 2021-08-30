package main

import (
	"fmt"
	"reflect"
)
type User struct{
	Id int
	Name string
	Age int
}

func (this *User) Call(){
	fmt.Println("user is called ..")
	fmt.Printf("%v\n",this)
}

func DoFiledAndMethod(input interface{}){
	inputType := reflect.TypeOf(input)
	inputValue := reflect.ValueOf(input)
	
	fmt.Println(inputType)
	fmt.Println(inputValue)

	for i:=0;i < inputType.NumField();i++{
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Println(field.Name,field.Type,value)
	}
}

func main(){
	var user User
	user = User{1,"tryer",10}
	DoFiledAndMethod(user)
}