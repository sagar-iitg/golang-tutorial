package main

import "fmt"

func main() {

	fmt.Println("Welcome to array")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Banana"
	fmt.Print("fruitList is", fruitList)
	array()

}
func array() {

	//compiler count the array elements for us
	b := [...]string{"Penn", "Teller"}
	fmt.Println()
	fmt.Println(b)

}
