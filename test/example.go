package main

import (
	"fmt"
	"strconv"
)

// If the number is divisible by 3, write "Foo" otherwise, the number
func Fooer(input int) string {

	isfoo := (input % 3) == 0
	fmt.Print(isfoo)

	if isfoo {
		return "Foo"
	}

	return strconv.Itoa(input)
}

func main() {
	Fooer(9)
}
