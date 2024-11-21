package main

/*
https://refactoring.guru/design-patterns/abstract-factory

Abstract Factory is a creational design pattern that lets you produce families of
related objects without specifying their concrete classes.

The first thing the Abstract Factory pattern suggests is to explicitly declare
interfaces for each distinct product of the product family (e.g., chair, sofa or coffee table).

Then you can make all variants of products follow those interfaces.
For example, all chair variants can implement the Chair interface;
all coffee table variants can implement the CoffeeTable interface, and so on.

The next move is to declare the Abstract Factory—an interface with a list of creation
methods for all products that are part of the product family
(for example, createChair, createSofa and createCoffeeTable)

These methods must return abstract product types represented by the interfaces we extracted previously:
Chair, Sofa, CoffeeTable and so on.

Now, how about the product variants? For each variant of a product family,
we create a separate factory class based on the AbstractFactory interface.
A factory is a class that returns products of a particular kind.
For example, the ModernFurnitureFactory can only create ModernChair, ModernSofa and ModernCoffeeTable objects.

The client code has to work with both factories and products via their respective abstract interfaces.
*/

// The abstract factory interface declares a set of methods that
// return different abstract products. These products are called
// a family and are related by a high-level theme or concept.
// Products of one family are usually able to collaborate among
// themselves. A family of products may have several variants,
// but the products of one variant are incompatible with the
// products of another variant.
type FurnitureFactory interface {
	createChair() Chair
	createSofa() Sofa
	createCoffeeTable() CoffeeTable
}

// Concrete factories produce a family of products that belong
// to a single variant. The factory guarantees that the
// resulting products are compatible. Signatures of the concrete
// factory's methods return an abstract product, while inside
// the method a concrete product is instantiated.
type ModernFurnitureFactory struct {
}

func (f *ModernFurnitureFactory) createChair() Chair {
	return &ModernChair{}
}

func (f *ModernFurnitureFactory) createSofa() Sofa {
	return &ModernSofa{}
}

func (f *ModernFurnitureFactory) createCoffeeTable() CoffeeTable {
	return &ModernCoffeeTable{}
}

type VictorianFurnitureFactory struct {
}

func (f *VictorianFurnitureFactory) createChair() Chair {
	return &VictorianChair{}
}

func (f *VictorianFurnitureFactory) createSofa() Sofa {
	return &VictorianSofa{}
}

func (f *VictorianFurnitureFactory) createCoffeeTable() CoffeeTable {
	return &VictorianCoffeeTable{}
}

// Each distinct product of a product family should have a base
// interface. All variants of the product must implement this
// interface.
type Chair interface {
	sitOn()
	setNumberOfLegs(numberOfLegs int)
}

type Sofa interface {
	sitOn()
	setNumberOfSeats(numberOfSeats int)
}

type CoffeeTable interface {
	setShape(shape string)
}
