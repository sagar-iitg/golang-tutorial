package main

import "math"

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {

	return math.Pi * c.radius * c.radius
}

func (r rectangle) perim() float64 {

	return 2*r.width + 2*r.height
}

func (c circle) perim() float64 {

	return math.Pi * c.radius * 2
}

func measure(g geometry) {

	println(g)
	println(g.area())
	println(g.perim())
}
func main() {

	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)

}
