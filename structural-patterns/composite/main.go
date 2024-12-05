package main

import (
	"fmt"
)

/*
https://refactoring.guru/design-patterns/composite

Composite is a structural design pattern that lets you compose objects into tree structures and then work with these structures
as if they were individual objects.

Using the Composite pattern makes sense only when the core model of your app can be represented as a tree.

For example, imagine that you have two types of objects: Products and Boxes. A Box can contain several
Products as well as a number of smaller Boxes.
These little Boxes can also hold some Products or even smaller Boxes, and so on.

Go over all the products and then calculate the total.
That would be doable in the real world; but in a program, it’s not as simple as running a loop.
You have to know the classes of Products and Boxes you’re going through,
the nesting level of the boxes and other nasty details beforehand.

The Composite pattern suggests that you work with Products and Boxes through a common interface
which declares a method for calculating the total price.

How would this method work?
For a product, it’d simply return the product’s price.
For a box, it’d go over each item the box contains, ask its price and then return a total for this box.
If one of these items were a smaller box, that box would also start going over its contents and so on,
until the prices of all inner components were calculated.
A box could even add some extra cost to the final price, such as packaging cost.

The greatest benefit of this approach is that you don’t need to care about the concrete classes of objects that compose the tree.
You don’t need to know whether an object is a simple product or a sophisticated box.
You can treat them all the same via the common interface.

Composite is a structural design pattern that lets you compose objects into tree structures
and then work with these structures as if they were individual objects.
*/

// The component interface declares common operations for both
// simple and complex objects of a composition.
type Graphic interface {
	draw()
}

// The leaf class represents end objects of a composition. A
// leaf object can't have any sub-objects. Usually, it's leaf
// objects that do the actual work, while composite objects only
// delegate to their sub-components.
type Dot struct {
	x, y int
}

func newDot(x, y int) *Dot {
	return &Dot{x, y}
}

func (d *Dot) draw() {
	fmt.Println("Dot at", d.x, d.y)
}

// All component classes can extend other components.
type Circle struct {
	radius int
}

func newCircle(radius int) *Circle {
	return &Circle{radius: radius}
}

func (c *Circle) draw() {
	fmt.Println("Circle with radius", c.radius)
}

// The composite class represents complex components that may
// have children. Composite objects usually delegate the actual
// work to their children and then "sum up" the result.
type CompoundGraphic struct {
	children []Graphic
}

func (c *CompoundGraphic) add(g Graphic) {
	c.children = append(c.children, g)
}

func (c *CompoundGraphic) remove(g Graphic) {
	for i, child := range c.children {
		if child == g {
			c.children = append(c.children[:i], c.children[i+1:]...)
		}
	}
}

// A composite executes its primary logic in a particular
// way. It traverses recursively through all its children,
// collecting and summing up their results. Since the
// composite's children pass these calls to their own
// children and so forth, the whole object tree is traversed
// as a result.
func (c *CompoundGraphic) draw() {
	for _, child := range c.children {
		child.draw()
	}
}

func main() {

	var all CompoundGraphic
	all.add(newDot(1, 2))
	all.add(newCircle(10))
	all.add(newCircle(20))
	all.draw()
}
