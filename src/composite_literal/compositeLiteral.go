package main

import "fmt"

type Car struct {
	make string
	model string
	year int
}

func (c *Car) Print() {
	fmt.Printf("\nThe car make is: %s, of model: %s, and built in: %d", c.make, c.model, c.year)
}

func main() {

	// Composite literal using default order
	audi := Car{"Audi", "A6", 2013}
	// Composite iternal using argument name (any order)
	bmw := Car{model: "M3", make: "BMW", year: 2014}
	// Use new keyword to create zeroed instance
	merc := new(Car)
	merc.model = "Bluetec"
	merc.make = "Mercedez-Benz"
	merc.year = 2015

	// Initialize an array using composite liternal

	cars := []*Car {&bmw, merc, &audi}

	bmw.Print()
	audi.Print()
	merc.Print()

	fmt.Println("\n==============================================")
	for _,cr := range cars {
		cr.Print()
	}
}