package main

/*
https://refactoring.guru/design-patterns/builder

Builder is a creational design pattern that lets you construct complex objects step by step.
The pattern allows you to produce different types and representations of an object using the same construction code.

Imagine a complex object that requires laborious, step-by-step initialization of many fields and
nested objects. Such initialization code is usually buried inside a monstrous constructor with lots
of parameters. Or even worse: scattered all over the client code.

The Builder pattern suggests that you extract the object construction
code out of its own class and move it to separate objects called builders.

The pattern organizes object construction into a set of steps (buildWalls, buildDoor, etc.).
To create an object, you execute a series of these steps on a builder object.
The important part is that you don’t need to call all of the steps. You can call only
those steps that are necessary for producing a particular configuration of an object.

You can go further and extract a series of calls to the builder steps you use to construct
a product into a separate class called director. The director class defines the
order in which to execute the building steps, while the builder provides the implementation for those steps.

Having a director class in your program isn’t strictly necessary.
You can always call the building steps in a specific order directly from the client code.
However, the director class might be a good place to put various construction routines so you can reuse them across your program.

Unlike other creational patterns, Builder doesn’t require products to have a common interface.
That makes it possible to produce different products using the same construction process.
*/

// The builder interface specifies methods for creating the
// different parts of the product objects.
type Builder interface {
	setSeats()
	setEngine()
	setTripComputer()
	setGPS()
	getCar() Car
}

func getBuilder(builderType string) Builder {
	switch builderType {
	case "suv":
		return newSUVBuilder()
	case "sport":
		return newSportBuilder()
	default:
		return nil
	}
}

// Product
// Using the Builder pattern makes sense only when your products
// are quite complex and require extensive configuration. The
// following two products are related, although they don't have
// a common interface.
type Car struct {
	seats        int
	engine       string
	tripComputer bool
	gps          bool
}

func (c *Car) String() {
	println("seats: ", c.seats)
	println("engine: ", c.engine)
	println("trip computer: ", c.tripComputer)
	println("gps: ", c.gps)
}

// Concrete Builder
// Unlike other creational patterns, builder lets you construct
// products that don't follow the common interface.
type SUVBuilder struct {
	seats        int
	engine       string
	tripComputer bool
	gps          bool
}

func newSUVBuilder() Builder {
	return &SUVBuilder{}
}

func (b *SUVBuilder) setSeats() {
	b.seats = 4
}

func (b *SUVBuilder) setEngine() {
	b.engine = "1.6"
}

func (b *SUVBuilder) setTripComputer() {
	b.tripComputer = true
}

func (b *SUVBuilder) setGPS() {
	b.gps = true
}

func (b *SUVBuilder) getCar() Car {
	return Car{
		seats:        b.seats,
		engine:       b.engine,
		tripComputer: b.tripComputer,
		gps:          b.gps,
	}
}

type SportBuilder struct {
	seats        int
	engine       string
	tripComputer bool
	gps          bool
}

func newSportBuilder() Builder {
	return &SportBuilder{}
}

func (b *SportBuilder) setSeats() {
	b.seats = 2
}

func (b *SportBuilder) setEngine() {
	b.engine = "3.0"
}

func (b *SportBuilder) setTripComputer() {
	b.tripComputer = false
}

func (b *SportBuilder) setGPS() {
	b.gps = false
}

func (b *SportBuilder) getCar() Car {
	return Car{
		seats:        b.seats,
		engine:       b.engine,
		tripComputer: b.tripComputer,
		gps:          b.gps,
	}
}

// The director is only responsible for executing the building
// steps in a particular sequence. It's helpful when producing
// products according to a specific order or configuration.
// Strictly speaking, the director class is optional, since the
// client can control builders directly.
type Director struct {
	builder Builder
}

func newDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) construct() Car {
	d.builder.setSeats()
	d.builder.setEngine()
	d.builder.setTripComputer()
	d.builder.setGPS()

	return d.builder.getCar()
}

func main() {
	builder := getBuilder("suv")

	director := newDirector(builder)

	car := director.construct()
	car.String()
}
