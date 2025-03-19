package main

/*
You are building a shape drawing application where different shapes can be drawn
(e.g., Circle, Rectangle, Triangle), and you want to add functionality to calculate
their areas, perimeters, and print information about them. Instead of modifying each
shape class to add new functionality, you will use the Visitor pattern to add
new operations without changing the existing Shape interfaces.
*/

import "fmt"

// Shape interface
type Shape interface {
	Accept(visitor Visitor)
}

// Visitor interface
type Visitor interface {
	VisitCircle(c *Circle)
	VisitRectangle(r *Rectangle)
	VisitTriangle(t *Triangle)
}

// Concrete Shape types

type Circle struct {
	radius float64
}

func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(c)
}

type Rectangle struct {
	width, height float64
}

func (r *Rectangle) Accept(visitor Visitor) {
	visitor.VisitRectangle(r)
}

type Triangle struct {
	base, height float64
}

func (t *Triangle) Accept(visitor Visitor) {
	visitor.VisitTriangle(t)
}

// Concrete Visitors

// AreaVisitor calculates the area of shapes
type AreaVisitor struct{}

func (v *AreaVisitor) VisitCircle(c *Circle) {
	area := 3.14159 * c.radius * c.radius
	fmt.Printf("Area of Circle: %.2f\n", area)
}

func (v *AreaVisitor) VisitRectangle(r *Rectangle) {
	area := r.width * r.height
	fmt.Printf("Area of Rectangle: %.2f\n", area)
}

func (v *AreaVisitor) VisitTriangle(t *Triangle) {
	area := 0.5 * t.base * t.height
	fmt.Printf("Area of Triangle: %.2f\n", area)
}

// PerimeterVisitor calculates the perimeter of shapes
type PerimeterVisitor struct{}

func (v *PerimeterVisitor) VisitCircle(c *Circle) {
	perimeter := 2 * 3.14159 * c.radius
	fmt.Printf("Perimeter of Circle: %.2f\n", perimeter)
}

func (v *PerimeterVisitor) VisitRectangle(r *Rectangle) {
	perimeter := 2 * (r.width + r.height)
	fmt.Printf("Perimeter of Rectangle: %.2f\n", perimeter)
}

func (v *PerimeterVisitor) VisitTriangle(t *Triangle) {
	// Assuming an equilateral triangle for simplicity
	perimeter := 3 * t.base
	fmt.Printf("Perimeter of Triangle: %.2f\n", perimeter)
}

// Main function demonstrating the Visitor pattern
func main() {
	// Create some shapes
	shapes := []Shape{
		&Circle{radius: 5},
		&Rectangle{width: 4, height: 6},
		&Triangle{base: 3, height: 4},
	}

	// Create visitors
	areaVisitor := &AreaVisitor{}
	perimeterVisitor := &PerimeterVisitor{}

	// Apply visitors to shapes
	for _, shape := range shapes {
		shape.Accept(areaVisitor)
		shape.Accept(perimeterVisitor)
	}
}
