package main

import "fmt"

/*
https://refactoring.guru/design-patterns/adapter

Adapter is a structural design pattern that allows objects with incompatible interfaces to collaborate.

You can create an adapter. This is a special object that converts the interface of one object so that another object can understand it.

An adapter wraps one of the objects to hide the complexity of conversion happening behind the scenes.
The wrapped object isn’t even aware of the adapter. For example, you can wrap an object that operates in meters and
kilometers with an adapter that converts all of the data to imperial units such as feet and miles.

Adapters can not only convert data into various formats but can also help objects with different interfaces collaborate.
Here’s how it works:

1. The adapter gets an interface, compatible with one of the existing objects.
2. Using this interface, the existing object can safely call the adapter’s methods.
3. Upon receiving a call, the adapter passes the request to the second object,
but in a format and order that the second object expects.
*/

/*
Suppose we have a legacy system where we interact with a Rectangle that calculates the area.
But we need to work with a Shape interface that only provides an Area() method.
We want to adapt the Rectangle class to conform to the Shape interface.
*/

type RectangleLegacy struct {
	width, height float64
}

func (r *RectangleLegacy) CalculateArea() float64 {
	return r.width * r.height
}

// New system
type Shape interface {
	getArea()
}

// Adapter Legacy to New
type RectangleAdapter struct {
	rectangle *RectangleLegacy
}

func (r *RectangleAdapter) getArea() {
	fmt.Println("Rectangle area: ", r.rectangle.CalculateArea())
}

func showArea(s Shape) {
	s.getArea()
}

type Circle struct {
	radius float64
}

func (c *Circle) getArea() {
	fmt.Println("Circle area: ", 3.14*c.radius*c.radius)
}

func main() {
	rectangle := &RectangleLegacy{width: 10, height: 5}
	rectangleAdapter := &RectangleAdapter{rectangle: rectangle}

	showArea(rectangleAdapter)

	circle := &Circle{radius: 10}
	showArea(circle)
}
