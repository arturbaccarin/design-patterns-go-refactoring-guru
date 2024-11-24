package main

import "fmt"

/*
https://refactoring.guru/design-patterns/abstract-factory

Abstract Factory is a creational design pattern that lets you produce families of
related objects without specifying their concrete classes.

You need a way to create individual furniture objects so that they match other objects of the same family.
Customers get quite mad when they receive non-matching furniture. (Think about a device interface, all buttons, pages and more need to be iOS or android)

Also, you don’t want to change existing code when adding new products or families of products to the program.

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
type GUIFactory interface {
	createButton() Button
	createHeader() Header
	createFooter() Footer
}

// Concrete factories produce a family of products that belong
// to a single variant. The factory guarantees that the
// resulting products are compatible. Signatures of the concrete
// factory's methods return an abstract product, while inside
// the method a concrete product is instantiated.
type WindowsGUIFactory struct {
}

func (w *WindowsGUIFactory) createButton() Button {
	return &WindowsButton{}
}

func (w *WindowsGUIFactory) createHeader() Header {
	return &WindowsHeader{}
}

func (w *WindowsGUIFactory) createFooter() Footer {
	return &WindowsFooter{}
}

type LinuxGUIFactory struct {
}

func (l *LinuxGUIFactory) createButton() Button {
	return &LinuxButton{}
}

func (l *LinuxGUIFactory) createHeader() Header {
	return &LinuxHeader{}
}

func (l *LinuxGUIFactory) createFooter() Footer {
	return &LinuxFooter{}
}

// Each distinct product of a product family should have a base
// interface. All variants of the product must implement this
// interface.
type Button interface {
	clickOn()
}

type Header interface {
	readImage()
}

type Footer interface {
	readNote()
}

// Concrete products are created by corresponding concrete
// factories.
type WindowsButton struct {
}

func (w *WindowsButton) clickOn() {
	fmt.Println("Clicking Windows button")
}

// Concrete product
type WindowsHeader struct {
}

func (w *WindowsHeader) readImage() {
	fmt.Println("Reading Windows image")
}

// Concrete product
type WindowsFooter struct {
}

func (w *WindowsFooter) readNote() {
	fmt.Println("Reading Windows note")
}

// Concrete product
type LinuxButton struct {
}

func (l *LinuxButton) clickOn() {
	fmt.Println("Clicking Linux button")
}

// Concrete product
type LinuxHeader struct {
}

func (l *LinuxHeader) readImage() {
	fmt.Println("Reading Linux image")
}

// Concrete product
type LinuxFooter struct {
}

func (l *LinuxFooter) readNote() {
	fmt.Println("Reading Linux note")
}

func GetGUIFactory(furniture string) (GUIFactory, error) {
	switch furniture {
	case "windows":
		return &WindowsGUIFactory{}, nil
	case "linux":
		return &LinuxGUIFactory{}, nil
	default:
		return nil, fmt.Errorf("unknown GUI type: %s", furniture)
	}
}

func main() {
	GUIFactory, err := GetGUIFactory("linux")
	if err != nil {
		panic(err)
	}

	button := GUIFactory.createButton()
	button.clickOn()

	header := GUIFactory.createHeader()
	header.readImage()

	footer := GUIFactory.createFooter()
	footer.readNote()
}
