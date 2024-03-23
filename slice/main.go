package main

import (
	"fmt"
	"sort"
)

func removeElement() {
	var fruitList = []string{"apple", "orange", "grape", "mango", "banana"}
	fmt.Println("Fruit List: ", fruitList)
	fruitList = append(fruitList[:2], fruitList[3:]...)
	fmt.Println("Fruit List: ", fruitList)

}
func main() {
	println("Wlocome to slices in Go!")
	var fruitList = []string{"apple", "orange", "grape", "mango", "banana"}
	fmt.Println("Fruit List: ", fruitList)
	fruitList = append(fruitList, "kiwi", "watermelon")
	fmt.Println("Fruit List: ", fruitList)
	fruitList = append(fruitList[2:3])
	fmt.Println("Fruit List: ", fruitList)
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 102
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 900w
	fmt.Println("High Scores: ", highScores)
	sort.Ints(highScores)
	fmt.Println("Sorted High Scores: ", highScores)
	fmt.Println("--------------------------------")
	removeElement()

}
