package main

import "fmt"

/*
You are designing a system to manage a set of shapes (like Circle and Square)
that can be drawn on different rendering devices (like Raster or Vector).
The challenge is to separate the abstraction (shape) from its implementation
(drawing method). Your goal is to apply the Bridge pattern to solve this problem.
*/

type Renderer interface {
	Render(shape Shape)
}

type Shape interface {
	Name() string
	Area() float64
}

type RasterRenderer struct {
}

func (r *RasterRenderer) Render(shape Shape) {
	fmt.Printf("Rasterizing %s with area %f\n", shape.Name(), shape.Area())
}

type VectorRenderer struct {
}

func (v *VectorRenderer) Render(shape Shape) {
	fmt.Printf("Vectorizing %s with area %f\n", shape.Name(), shape.Area())
}

type Circle struct {
	name   string
	radius float64
}

func (c *Circle) Name() string {
	return c.name
}

func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

type Square struct {
	name string
	side float64
}

func (s *Square) Name() string {
	return s.name
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

func main() {
	var renderer Renderer
	renderer = &RasterRenderer{}

	var shape Shape
	shape = &Circle{name: "Circle", radius: 5}
	renderer.Render(shape)

	renderer = &VectorRenderer{}

	shape = &Square{name: "Square", side: 5}
	renderer.Render(shape)
}
