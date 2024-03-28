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
	s.GetStatus()

}

func (s student) GetStatus() {
	fmt.Println("Student is active")
	fmt.Println(s.name, "is active")
	fmt.Println(s.age, "is active")

}
