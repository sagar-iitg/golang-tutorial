package main

import "fmt"

type A struct {
	year int
}

func (a A) Greet() {
	fmt.Println("Hello Golang", a.year)
}

type B struct {
	A
}

func (b B) Greet() {
	fmt.Println("Welcome Golang", b.year)
}

func main() {

	var a A
	a.year = 2021
	var b B
	b.year = 2021
	a.Greet()
	b.Greet()

}
