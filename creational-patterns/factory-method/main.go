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
type DeliveryService interface {
	Deliver()
	setName(name string)
	setTime(time int)
	getName() string
	getTime() int
}

// Concrete product: Concrete Products are different implementations of the product interface.
type DeliveryType struct {
	name string
	time int
}

func (d *DeliveryType) setName(name string) {
	d.name = name
}

func (d *DeliveryType) getName() string {
	return d.name
}

func (d *DeliveryType) setTime(time int) {
	d.time = time
}

func (d *DeliveryType) getTime() int {
	return d.time
}

func (d *DeliveryType) Deliver() {
	fmt.Printf("Delivery %s in %d hours\n", d.name, d.time)
}

// Concrete Creators override the base factory method so it returns a different type of product.
type Ship struct {
	DeliveryType
}

func newShip() DeliveryService {
	return &Ship{
		DeliveryType: DeliveryType{
			name: "Ship",
			time: 3,
		},
	}
}

// Concrete Creators
type Airplane struct {
	DeliveryType
}

func newAirplane() DeliveryService {
	return &Airplane{
		DeliveryType: DeliveryType{
			name: "Airplane",
			time: 1,
		},
	}
}

// Factory
func createDeliveryService(deliveryType string) DeliveryService {
	switch deliveryType {
	case "ship":
		return newShip()
	case "airplane":
		return newAirplane()
	default:
		return nil
	}
}

func main() {
	transport := createDeliveryService("airplane")
	transport.Deliver()
}

/*
Applicability
Use the Factory Method when you don’t know beforehand the exact types and dependencies of the objects your code should work with.

The Factory Method separates product construction code from the code that actually uses the product.
Therefore it’s easier to extend the product construction code independently from the rest of the code.

To add a new product type to the app, you’ll only need to create a new creator subclass and override the factory method in it.
*/
