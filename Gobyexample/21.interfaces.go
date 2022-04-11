package main

import(
	"fmt"
	"math"
)

/*
	Interfaces are named collections of method signatures.
	Here's a basic interface for geometric shapes.
	For our example we'll implement this interface on rect and circle types.
	To implement an interface in go, we just need to implement
	all the methods in the interface. Here we implement geometry on rects.
	The implementation for circles.
	If a variable has an interface type, then we can call methods that are 
	in the named interface. Here's a generic measure function taking advantage 
	of this to work on any geometry.

	The circle and rect struct types both implement the geometry interface 
	so we can use instances of these structs as arguments to measure.

	To know more about go's great interfaces check out this post: 
	https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
*/

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2 * r.width + 2 * r.height
}

func (c circle) area() float64 {
	return math.pi * c.radius * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}