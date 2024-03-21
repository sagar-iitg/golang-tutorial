package main

import "fmt"

func pointer() {
	var x *int
	fmt.Println("Value of pointer is: ", x)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of actual data is: ", *ptr)
	fmt.Println("Value of pointer is: ", ptr)
	*ptr = *ptr + 5
	fmt.Print(myNumber)

}
func main() {
	var x int = 10
	p := &x
	println(x)
	println(*p)
	pointer()

}
