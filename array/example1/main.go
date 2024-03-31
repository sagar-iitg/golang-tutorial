package main

import "fmt"

func main() {

	array := [...]float64{67.7, 89.8, 21, 78.9}
	x := Sum(&array)
	fmt.Println(x)
}
func Sum(a *[4]float64) (sum float64) {
	for _, v := range a {
		fmt.Println(v)
		sum += v
	}
	return sum
}
