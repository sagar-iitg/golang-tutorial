package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is one")
	case 2:
		fmt.Println("Dice value is two")
	case 3:
		fmt.Println("Dice value is three")
	case 4:
		fmt.Println("Dice value is four")
		fallthrough
	case 5:
		fmt.Println("Dice Value is five")
	default:
		fmt.Println("What was that!")
	}

}
