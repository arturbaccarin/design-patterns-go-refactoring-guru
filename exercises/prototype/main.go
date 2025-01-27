package main

import "fmt"

/*
In this exercise, you'll implement the Prototype Pattern by creating
a Vehicle interface and implementing it for two types of vehicles: Car and Truck.
The goal is to clone a Car or Truck object and modify some properties of the clone,
rather than creating new instances from scratch.
*/

type Vehicle interface {
	Clone() Vehicle
}

type Car struct {
	Model string
	Brand string
	Year  int
	Color string
}

func NewCar(model, brand string, year int) Car {
	return Car{
		Model: model,
		Brand: brand,
		Year:  year,
	}
}

func (c *Car) setColor(color string) {
	c.Color = color
}

func (c Car) Clone() Vehicle {
	return &Car{
		Model: c.Model,
		Brand: c.Brand,
		Year:  c.Year,
		Color: c.Color,
	}
}

type Truck struct {
	Model        string
	Brand        string
	Year         int
	LoadCapacity int
}

func NewTruck(model, brand string, year int, loadCapacity int) *Truck {
	return &Truck{
		Model:        model,
		Brand:        brand,
		Year:         year,
		LoadCapacity: loadCapacity,
	}
}

func (t *Truck) setLoadCapacity(capacity int) {
	t.LoadCapacity = capacity
}

func (t *Truck) Clone() Vehicle {
	return &Truck{
		Model:        t.Model,
		Brand:        t.Brand,
		Year:         t.Year,
		LoadCapacity: t.LoadCapacity,
	}
}

func main() {
	var vehicle Vehicle

	vehicle = NewCar("Civic", "Honda", 2022)

	vehicleClone := vehicle.Clone()
	vehicleClone.(*Car).setColor("Red")

	fmt.Println(vehicleClone.(*Car).Color)
}
