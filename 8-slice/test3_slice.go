package main

import "fmt"

func main() {
	var numbers = make([]int, 3)
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 1, 2, 3, 4, 5, 6)
	numbers = append(numbers, 7, 8)
	fmt.Printf("len=%d,cap=%d,slice=%v \n", len(numbers), cap(numbers), numbers)
	s1 := make([]int, 3)
	copy(s1, numbers[0:2])
	fmt.Printf("%v", s1)

}
