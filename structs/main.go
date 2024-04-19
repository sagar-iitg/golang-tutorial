package main

import (
	"fmt"
	"time"
)

type student struct {
	name string
	age  int
}

type user struct {
	firstName string
	lastName  string
	birthDate string
	createAt  time.Time
}

func main() {
	s := student{name: "John", age: 25}
	fmt.Println(s)
	fmt.Printf("student details are %+v\n", s)

	appUser := user{
		firstName: "sagar",
		lastName:  "kumar",
		birthDate: "1995-12-12",
		createAt:  time.Now(),
	}
	outputUserDetail(appUser)

}

func outputUserDetail(u user) {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createAt)
}
