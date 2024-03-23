package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	s := student{name: "John", age: 25}
	fmt.Println(s)
	fmt.Printf("student details are %+v\n", s)

}
