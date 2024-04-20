package main

import "fmt"

func main() {
	age := 32
	fmt.Println(age)
	fmt.Print(getAdultYear(age))
}
func getAdultYear(age int) int {
	return age - 18
}
