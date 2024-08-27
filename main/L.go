//Sem usar (Liskov Substitution Principle - LSP)
// package main
//
// import "fmt"
//
// type Bird struct{}
//
//	func (b *Bird) Fly() {
//		fmt.Println("Flying")
//	}
//
//	type Ostrich struct {
//		Bird
//	}
//
//	func main() {
//		ostrich := Ostrich{}
//		ostrich.Fly() // Ostriches can't fly, this violates LSP
//	}

// Usando (Liskov Substitution Principle - LSP)
package main

import "fmt"

// Bird Classe base
type Bird interface {
	Move() string
}

// Sparrow Pardal: voa
type Sparrow struct{}

func (s Sparrow) Move() string {
	return s.Fly()
}

func (s Sparrow) Fly() string {
	return "Sparrow is flying"
}

// Ostrich Avestruz: corre
type Ostrich struct{}

func (o Ostrich) Move() string {
	return o.Run()
}

func (o Ostrich) Run() string {
	return "Ostrich is running"
}

// Penguin Pinguim: nada
type Penguin struct{}

func (p Penguin) Move() string {
	return p.Swim()
}

func (p Penguin) Swim() string {
	return "Penguin is swimming"
}

func main() {
	birds := []Bird{
		Sparrow{},
		Ostrich{},
		Penguin{},
	}
	for _, bird := range birds {
		fmt.Println(bird.Move())
	}
}
