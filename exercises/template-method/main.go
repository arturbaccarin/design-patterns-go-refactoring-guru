package main

/*
In this exercise, you'll simulate the process of cooking two types of meals: Tea and Coffee. The cooking process for both meals is largely the same, but there are certain differences in how the final product is prepared (e.g., boiling water for tea vs. making coffee).

Steps for the cooking process:
Boil water.
Add specific ingredients (tea leaves or coffee grounds).
Brew the beverage.
Serve the beverage.
The TemplateMethod will define the basic steps, and the subclasses will override specific parts of the process.
*/

import "fmt"

// Beverage is the abstract base class that defines the template method
type Beverage interface {
	Cook()           // Template method that defines the overall process
	BoilWater()      // Step 1: Boil water
	AddIngredients() // Step 2: Add specific ingredients
	Brew()           // Step 3: Brew the beverage
	Serve()          // Step 4: Serve the beverage
}

// Abstract implementation of Beverage (common steps)
type AbstractBeverage struct{}

func (ab *AbstractBeverage) BoilWater() {
	fmt.Println("Boiling water...")
}

func (ab *AbstractBeverage) Serve() {
	fmt.Println("Serving the beverage.")
}

// Tea class implements the Beverage interface and overrides specific methods
type Tea struct {
	*AbstractBeverage // Embedding common methods
}

func (t *Tea) Cook() {
	t.BoilWater()
	t.AddIngredients()
	t.Brew()
	t.Serve()
}

func (t *Tea) AddIngredients() {
	fmt.Println("Adding tea leaves...")
}

func (t *Tea) Brew() {
	fmt.Println("Brewing the tea...")
}

// Coffee class implements the Beverage interface and overrides specific methods
type Coffee struct {
	*AbstractBeverage // Embedding common methods
}

func (c *Coffee) Cook() {
	c.BoilWater()
	c.AddIngredients()
	c.Brew()
	c.Serve()
}

func (c *Coffee) AddIngredients() {
	fmt.Println("Adding coffee grounds...")
}

func (c *Coffee) Brew() {
	fmt.Println("Brewing the coffee...")
}

func main() {
	// Create a Tea object and cook it
	tea := &Tea{}
	tea.Cook()

	// Create a Coffee object and cook it
	coffee := &Coffee{}
	coffee.Cook()
}
