package main

import "fmt"

/*
https://refactoring.guru/design-patterns/factory-method

Factory Method is a creational design pattern that provides an interface for creating objects in a superclass,
but allows subclasses to alter the type of objects that will be created.

Factory method is a creational design pattern which solves the problem of creating product objects without specifying their concrete classes.

The Factory Method pattern suggests that you replace direct object construction calls (using the new operator) with calls to a special factory method.

Don’t worry: the objects are still created via the new operator, but it’s being called from within the factory method. Objects returned by a factory method are often referred to as products.

Subclasses may return different types of products only if these products have a common base class or interface.
Also, the factory method in the base class should have its return type declared as this interface.

It’s impossible to implement the classic Factory Method pattern in Go due to lack of OOP features such as classes and inheritance.
However, we can still implement the basic version of the pattern, the Simple Factory.
*/

// Product interface: declares the interface, which is common to all objects that can be produced by the creator and its subclasses.
type Button interface {
	render()
	onClick()
}

// Concrete product: Concrete Products are different implementations of the product interface.
type WindowsButton struct {
}

func (b *WindowsButton) render() {
	fmt.Println("Rendering Windows button")
}

func (b *WindowsButton) onClick() {
	fmt.Println("Clicking Windows button")
}

type HTMLButton struct {
}

func (b *HTMLButton) render() {
	fmt.Println("Rendering HTML button")
}

func (b *HTMLButton) onClick() {
	fmt.Println("Clicking HTML button")
}

// Concrete Creators override the base factory method so it returns a different type of product.
func newWindowsButton() Button {
	return &WindowsButton{}
}

func newHTMLButton() Button {
	return &HTMLButton{}
}

// Factory
func createButton(system string) Button {
	switch system {
	case "windows":
		return newWindowsButton()
	case "web":
		return newHTMLButton()
	default:
		return nil
	}
}

func main() {
	button := createButton("web")
	button.render()
	button.onClick()
}

/*
Applicability
Use the Factory Method when you don’t know beforehand the exact types and dependencies of the objects your code should work with.

The Factory Method separates product construction code from the code that actually uses the product.
Therefore it’s easier to extend the product construction code independently from the rest of the code.

To add a new product type to the app, you’ll only need to create a new creator subclass and override the factory method in it.
*/
