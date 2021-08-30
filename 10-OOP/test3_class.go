package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) SetName(name string) {
	this.name = name
}

func (this *Human) Eat() {
	fmt.Println("eating...")
}

type SuperMan struct {
	Human
	level int
}

func (this *SuperMan) fly() {
	fmt.Println("flying...")
}
func (this *SuperMan) Print() {
	fmt.Println(this.name)
	fmt.Println(this.sex)
	fmt.Println(this.level)
}
func main() {
	man := Human{"zs", "female"}
	fmt.Println(man)
	sp := SuperMan{Human{"sp", "male"}, 10}
	fmt.Println(sp)
	sp.fly()
	sp.Print()
}
