package main

import "fmt"

func main() {
	//var myMap map[string]string
	/* if myMap == nil {
		fmt.Println("myMap is nil")
	} */
	myMap := make(map[string]string, 2)
	myMap["one"] = "one"
	myMap["two"] = "two"
	myMap["three"] = "three"
	fmt.Printf("%v\n", myMap)

	myMap3 := map[string]string{"one": "one", "two": "two", "three": "three"}
	fmt.Println(myMap3)
}
