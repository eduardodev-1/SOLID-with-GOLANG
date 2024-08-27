//Sem Dependency Inversion Principle (DIP)
// package main
//
// import "fmt"
//
// type LightBulb struct{}
//
//	func (lb *LightBulb) TurnOn() {
//		fmt.Println("Lightbulb on")
//	}
//
//	type Switch struct {
//		bulb LightBulb
//	}
//
//	func (s *Switch) Operate() {
//		s.bulb.TurnOn()
//	}
//
//	func main() {
//		lightSwitch := Switch{bulb: LightBulb{}}
//		lightSwitch.Operate()
//	}

// Usando Dependency Inversion Principle (DIP)
package main

import "fmt"

type Switchable interface {
	TurnOn()
}

type LightBulb struct{}

func (lb *LightBulb) TurnOn() {
	fmt.Println("Lightbulb on")
}

type Fan struct{}

func (f *Fan) TurnOn() {
	fmt.Println("Fan on")
}

type Switch struct {
	device Switchable
}

func (s *Switch) Operate() {
	s.device.TurnOn()
}

func main() {
	lightBulb := LightBulb{}
	fan := Fan{}

	lightSwitch := Switch{device: &lightBulb}
	lightSwitch.Operate() // Operando a lâmpada

	fanSwitch := Switch{device: &fan}
	fanSwitch.Operate() // Operando o ventilador
}
