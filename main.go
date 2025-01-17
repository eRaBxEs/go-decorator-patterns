package main

import (
	"fmt"
)

// Component interface
type Coffee interface {
	Cost() float64
	Description() string
}

// Concrete Component
type SimpleCoffee struct{}

func (s *SimpleCoffee) Cost() float64 {
	return 2.0
}

func (s *SimpleCoffee) Description() string {
	return "Simple Coffee"
}

// Decorator 1: Milk
type Milk struct { // here the milk decorator wraps around another coffee object
	coffee Coffee
}

func (m *Milk) Cost() float64 {
	return m.coffee.Cost() + 0.5
}

func (m *Milk) Description() string {
	return m.coffee.Description() + ", Milk"
}

// Decorator 2: Caramel
type Caramel struct { // here the caramel decorartor wraps around another coffee object
	coffee Coffee
}

func (c *Caramel) Cost() float64 {
	return c.coffee.Cost() + 1.0
}

func (c *Caramel) Description() string {
	return c.coffee.Description() + ", Caramel"
}

func main() {

	// create a simple coffee
	coffee := &SimpleCoffee{}

	// Decorate it with Milk
	coffeeWithMilk := &Milk{coffee: coffee}
	// output
	fmt.Println("Coffee with Milk:", coffeeWithMilk.Description(), "- Cost:", coffeeWithMilk.Cost())

	// Decorate it with Caramel
	coffeeWithCaramel := &Caramel{coffee: coffee}
	fmt.Println("Coffee with Caramel:", coffeeWithCaramel.Description(), "- Cost:", coffeeWithCaramel.Cost())

	// combining multiple decorators
	coffeeWithMilkAndCaramel := &Milk{coffee: &Caramel{coffee: coffee}}
	fmt.Println("Coffee with Milk and Caramel:", coffeeWithMilkAndCaramel.Description(), "- Cost:", coffeeWithMilkAndCaramel.Cost())

	// another multiple decorator combination, just to see if caramel comes before milk string wise
	coffeeWIthCaramelAndMilk := &Caramel{coffee: &Milk{coffee: coffee}}
	fmt.Println("Coffee with Milk and Caramel:", coffeeWIthCaramelAndMilk.Description(), "- Cost:", coffeeWIthCaramelAndMilk.Cost())

}
