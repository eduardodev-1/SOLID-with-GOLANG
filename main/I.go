//Sem usar (Interface Segregation Principle - ISP)
// package main
//
// import "fmt"
//
//	type Worker interface {
//		Work()
//		Eat()
//	}
//
// type Human struct{}
//
//	func (h Human) Work() {
//		fmt.Println("Human working")
//	}
//
//	func (h Human) Eat() {
//		fmt.Println("Human eating")
//	}
//
// type Robot struct{}
//
//	func (r Robot) Work() {
//		fmt.Println("Robot working")
//	}
//
//	func (r Robot) Eat() {
//		// Robots can't eat, but still need to implement this
//	}

// Usando (Interface Segregation Principle - ISP)
package main

import "fmt"

type Worker interface {
	Work()
}

type Eater interface {
	Eat()
}

type Human struct{}

func (h Human) Work() {
	fmt.Println("Human working")
}

func (h Human) Eat() {
	fmt.Println("Human eating")
}

type Robot struct{}

func (r Robot) Work() {
	fmt.Println("Robot working")
}

func main() {
	var worker Worker = Human{}
	worker.Work()

	var eater Eater = Human{}
	eater.Eat()

	robot := Robot{}
	robot.Work()
}
