// Sem usar (Open/Closed Principle - OCP)
// package main
//
// import "fmt"
//
//	type Rectangle struct {
//		Width, Height float64
//	}
//
//	type Circle struct {
//		Radius float64
//	}
//
//	func Area(shape interface{}) float64 {
//		switch shape.(type) {
//		case Rectangle:
//			r := shape.(Rectangle)
//			return r.Width * r.Height
//		case Circle:
//			c := shape.(Circle)
//			return 3.14 * c.Radius * c.Radius
//		}
//		return 0
//	}
//
//	func main() {
//		rect := Rectangle{10, 5}
//		circ := Circle{7}
//		fmt.Println("Rectangle Area:", Area(rect))
//		fmt.Println("Circle Area:", Area(circ))
//	}

// Usando (Open/Closed Principle - OCP)
package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	shapes := []Shape{
		Rectangle{10, 5},
		Circle{7},
	}

	for _, shape := range shapes {
		fmt.Println("Area:", shape.Area())
	}
}
