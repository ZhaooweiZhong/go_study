package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this Hero) GetName() {
	fmt.Println("GetName: ", this.Name)
}

func (this Hero) GetLevel() int {
	return this.Level
}
func (zzw *Hero) SetLevel() {
	zzw.Level = 12
}
func main() {
	var hero Hero
	hero.Name = "hero"
	hero.Ad = 100
	hero.Level = 10
	hero.GetName()
	fmt.Println(hero.GetLevel())
	hero.SetLevel()
	fmt.Println(hero)
}
