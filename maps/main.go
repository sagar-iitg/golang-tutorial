package main

import "fmt"

func main() {

	fmt.Println("Maps in Go")
	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	myMap["four"] = 4
	fmt.Println(myMap)
	fmt.Println(myMap["three"])
	delete(myMap, "three")
	fmt.Println(myMap)

	//loops are intersting

	for key, value := range myMap {
		fmt.Println(key, value)
	}

}
