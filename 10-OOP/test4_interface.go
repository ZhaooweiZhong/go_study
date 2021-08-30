package main

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}
type Cat struct {
	Color string
	Type  string
}

func (this *Cat) GetColor() string {
	return this.Color
}

func (this *Cat) GetType() string {
	return this.Type
}

func (this *Cat) Sleep() {
	fmt.Println("zzz......")
}

func PrintAnimal(animal AnimalIF) {
	fmt.Println(animal.GetType())
}

func main() {
	var animal AnimalIF
	animal = &Cat{"cat", "white"}
	animal.Sleep()
	PrintAnimal(&Cat{"cat", "white"})
}
