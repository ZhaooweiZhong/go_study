package test

import "fmt"

func printArray(myArray [4]int) {
	for index, i := range myArray {
		fmt.Println("index:", index, "value:", i)
	}
}

func test() {
	myArray := [4]int{1, 2, 3, 4}
	printArray(myArray)
}
