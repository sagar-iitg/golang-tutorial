package main

import "fmt"

// taking a structure
type Employee struct {

	// taking variables
	name  string
	empid int
}

func main() {
	emp := Employee{"ABC", 19078}
	// Here, it is the pointer to the struct
	pts := &emp
	fmt.Println(pts)
	fmt.Println(pts.name)
	fmt.Println((*pts).name)

}
