package main

import (
	"fmt"
	"sync"
)

/*
You are developing a game world simulation where you need to manage various objects like trees,
rocks, and animals in a forest. Each of these objects may have shared characteristics (e.g., type, color),
but also some unique properties (e.g., position, size, specific attributes like animal age).

The goal is to optimize memory usage by sharing common data between objects of the same type
(e.g., all Oak trees should share the same species and color).
*/

// Flyweight
type GameObjectType struct {
	Name    string
	Color   string
	Species string
}

type GameObject interface {
	Render()
}

type Tree struct {
	X, Y     int
	TreeType *GameObjectType
}

func (t *Tree) Render() {
	fmt.Printf("Tree of species %s at position (%d, %d) with color %s\n", t.TreeType.Species, t.X, t.Y, t.TreeType.Color)
}

type Animal struct {
	X, Y       int
	AnimalType *GameObjectType
	Age        int
}

func (a *Animal) Render() {
	fmt.Printf("Animal of type %s at position (%d, %d), age: %d\n", a.AnimalType.Species, a.X, a.Y, a.Age)
}

type GameObjectFactory struct {
	objectTypes map[string]*GameObjectType
	mu          sync.Mutex
}

func NewGameObjectFactory() *GameObjectFactory {
	return &GameObjectFactory{
		objectTypes: make(map[string]*GameObjectType),
	}
}

func (f *GameObjectFactory) GetGameObjectType(name, color, species string) *GameObjectType {
	f.mu.Lock()
	defer f.mu.Unlock()

	key := name + "-" + species + "-" + color
	if object, exists := f.objectTypes[key]; exists {
		return object
	}

	newObject := &GameObjectType{Name: name, Color: color, Species: species}
	f.objectTypes[key] = newObject
	return newObject
}

func main() {
	factory := NewGameObjectFactory()

	objects := []GameObject{
		&Tree{X: 5, Y: 10, TreeType: factory.GetGameObjectType("Tree", "Green", "Oak")},
		&Tree{X: 10, Y: 15, TreeType: factory.GetGameObjectType("Tree", "Green", "Oak")},
		&Tree{X: 20, Y: 25, TreeType: factory.GetGameObjectType("Tree", "Brown", "Pine")},
		&Animal{X: 30, Y: 35, AnimalType: factory.GetGameObjectType("Animal", "Brown", "Deer"), Age: 5},
		&Animal{X: 50, Y: 50, AnimalType: factory.GetGameObjectType("Animal", "Gray", "Wolf"), Age: 3},
		&Animal{X: 60, Y: 70, AnimalType: factory.GetGameObjectType("Animal", "Gray", "Wolf"), Age: 1},
	}

	for _, object := range objects {
		object.Render()
	}
}
