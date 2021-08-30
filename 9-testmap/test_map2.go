package main

import "fmt"

func printMap(mymap map[string]string) {
	fmt.Println(mymap)
	mymap["USA"] = "WC."
}

func main() {
	ci := map[string]string{
		"China": "BEIJING",
		"Japan": "Tokyo",
		"USA":   "NewYork",
	}
	for _, value := range ci {
		fmt.Println("value:", value)
	}
	printMap(ci)
	delete(ci, "China")
	for _, value := range ci {
		fmt.Println("value:", value)
	}
}
