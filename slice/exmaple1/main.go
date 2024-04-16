package main

import "fmt"

//function to modify slice
func modifySlice(numbers *[]int) {
	//modify the slice

	for i := range *numbers {
		(*numbers)[i] *= 2
	}

}

func main() {

	//create slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Original Slice: ", numbers)
	modifySlice(&numbers)
	fmt.Println("Modified Slice: ", numbers)

}
