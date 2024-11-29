package main

import "fmt"

/*
https://refactoring.guru/design-patterns/prototype

Prototype is a creational design pattern that lets you copy existing objects without making your code dependent on their classes.

Say you have an object, and you want to create an exact copy of it.
How would you do it? First, you have to create a new object of the same class.
Then you have to go through all the fields of the original object and copy their values over to the new object.

There’s one more problem with the direct approach.
Since you have to know the object’s class to create a duplicate, your code becomes dependent on that class.

Sometimes you only know the interface that the object follows, but not its concrete class, when,
for example, a parameter in a method accepts any objects that follow some interface.

The Prototype pattern delegates the cloning process to the actual objects that are being cloned.
The pattern declares a common interface for all objects that support cloning.
This interface lets you clone an object without coupling your code to the class of that object.

The implementation of the clone method is very similar in all classes.
The method creates an object of the current class and carries over all of the field values of the old object into the new one.
You can even copy private fields because most programming languages let objects access private fields of other objects that belong
to the same class.

An object that supports cloning is called a prototype.
When your objects have dozens of fields and hundreds of possible configurations, cloning them might serve as an alternative to subclassing.
*/

// Base prototype
type Shape interface {
	getArea()
	clone() Shape
}

// Concrete prototype. The cloning method creates a new object
// in one go by calling the constructor of the current class and
// passing the current object as the constructor's argument.
// Performing all the actual copying in the constructor helps to
// keep the result consistent: the constructor will not return a
// result until the new object is fully built; thus, no object
// can have a reference to a partially-built clone.
type Rectangle struct {
	width, height float64
}

func newRectangle(width, height float64) *Rectangle {
	return &Rectangle{width: width, height: height}
}

func (r *Rectangle) getArea() {
	fmt.Println("Rectangle area: ", r.width*r.height)
}

func (r *Rectangle) clone() Shape {
	return &Rectangle{width: r.width, height: r.height}
}

type Circle struct {
	radius float64
}

func newCircle(radius float64) *Circle {
	return &Circle{radius: radius}
}

func (c *Circle) getArea() {
	fmt.Println("Circle area: ", 3.14*c.radius*c.radius)
}

func (c *Circle) clone() Shape {
	return &Circle{radius: c.radius}
}

func main() {
	c := newCircle(10)
	c.getArea()

	c2 := c.clone()
	c2.getArea()

	c3 := c2.(*Circle)
	c3.radius = 20

	c3.getArea()
}
