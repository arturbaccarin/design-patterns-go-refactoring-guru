package main

import "fmt"

/*
https://refactoring.guru/design-patterns/visitor

Visitor is a behavioral design pattern that lets you separate
algorithms from the objects on which they operate.

Imagine that your team develops an app which works with
geographic information structured as one colossal graph.
Each node of the graph may represent a complex entity such as a city,
but also more granular things like industries, sightseeing areas, etc.

The nodes are connected with others if there’s a road between
the real objects that they represent. Under the hood, each node
type is represented by its own class, while each specific node is an object.

At some point, you got a task to implement exporting the graph into XML format.
At first, the job seemed pretty straightforward. You planned to add an export
method to each node class and then leverage recursion to go over each node of
the graph, executing the export method. The solution was simple and elegant:
thanks to polymorphism, you weren’t coupling the code which called the export
method to concrete classes of nodes.

At some point, you got a task to implement exporting the graph into XML format.
At first, the job seemed pretty straightforward. You planned to add an export
method to each node class and then leverage recursion to go over each node of the graph,
executing the export method. The solution was simple and elegant: thanks to polymorphism,
you weren’t coupling the code which called the export method to concrete classes of nodes.

The Visitor pattern suggests that you place the new behavior into a
separate class called visitor, instead of trying to integrate it into existing classes.
The original object that had to perform the behavior is now passed to one of the
visitor’s methods as an argument, providing the method access to all necessary
data contained within the object.

Visitor is a behavioral design pattern that allows adding new behaviors to
existing class hierarchy without altering any existing code.
*/

/*
Conceptual Example

The Visitor pattern lets you add behavior to a struct without actually modifying the
struct. Let’s say you are the maintainer of a lib which has
different shape structs such as:

Square
Circle
Triangle

Each of the above shape structs implements the common shape interface.

A team requested you to add the getArea behavior to the shape structs.

The functions visitForSquare(square), visitForCircle(circle),
visitForTriangle(triangle) will let us add functionality to squares,
circles and triangles respectively.

Wondering why can’t we have a single method visit(shape) in the visitor interface?
The reason is that the Go language doesn’t support method overloading, so you can’t
have methods with the same names but different parameters
*/

// Element

type Shape interface {
	getType() string
	accept(Visitor)
}

// Concrete Element
type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

// Concrete Element
type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

type Rectangle struct {
	l int
	b int
}

func (t *Rectangle) accept(v Visitor) {
	v.visitForrectangle(t)
}

func (t *Rectangle) getType() string {
	return "rectangle"
}

// Visitor
type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForrectangle(*Rectangle)
}

// Concrete Visitor
type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	// Calculate area for square.
	// Then assign in to the area instance variable.
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
}
func (a *AreaCalculator) visitForrectangle(s *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
	// Calculate middle point coordinates for square.
	// Then assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *MiddleCoordinates) visitForrectangle(t *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}

func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
