// package main

// import "fmt"

// func main() {
// 	numbers := []int{1, 2, 3}
// 	modifySlice(numbers)
// 	fmt.Println("Modified Slice (Passed by Value):", numbers)
// }

// // Function to modify a slice.
// func modifySlice(s []int) {

// 	for i := range s {
// 		s[i] *= 2
// 	}
// }

package main

import "fmt"

// Function to modify a slice.
func modifySlice(s []int) {
	for i := range s {
		s[i] *= 2
	}
}

func main() {
	numbers := []int{1, 2, 3}
	modifySlice(numbers)
	fmt.Println("Original Slice:", numbers) // Output: [1 2 3]
}
