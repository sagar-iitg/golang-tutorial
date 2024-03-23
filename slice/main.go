package main

import "fmt"

func main() {
	println("Wlocome to slices in Go!")
	var fruitList = []string{"apple", "orange", "grape", "mango", "banana"}
	fmt.Println("Fruit List: ", fruitList)
	fruitList = append(fruitList, "kiwi", "watermelon")
	fmt.Println("Fruit List: ", fruitList)
	fruitList = append(fruitList[2:3])
	fmt.Println("Fruit List: ", fruitList)

}
