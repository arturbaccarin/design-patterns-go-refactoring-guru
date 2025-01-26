package main

import "fmt"

/*
You're building a game that has multiple characters and weapons.
These characters belong to different factions, and each faction
has its own style for characters and weapons. The goal is to create
a system where each faction produces its own set of characters and weapons
without having to modify the client code each time you add a new faction.
*/
type Bow struct {
	name string
}

func (b *Bow) attack() {
	fmt.Printf("Attacking with %s bow\n", b.name)
}

type Sword struct {
	name string
}

func (s *Sword) attack() {
	fmt.Printf("Attacking with %s sword\n", s.name)
}

type Axe struct {
	name string
}

func (a *Axe) attack() {
	fmt.Printf("Attacking with %s axe\n", a.name)
}

type AbstractFactory interface {
	createBow() Bow
	createSword() Sword
	createAxe() Axe
}

type ElfFactory struct {
}

func (ElfFactory) createBow() Bow {
	return Bow{
		name: "Elf",
	}
}

func (ElfFactory) createSword() Sword {
	return Sword{
		name: "Elf",
	}
}

func (ElfFactory) createAxe() Axe {
	return Axe{
		name: "Elf",
	}
}

type OrcFactory struct {
}

func (OrcFactory) createBow() Bow {
	return Bow{
		name: "Orc",
	}
}

func (OrcFactory) createSword() Sword {
	return Sword{
		name: "Orc",
	}
}

func (OrcFactory) createAxe() Axe {
	return Axe{
		name: "Orc",
	}
}

func getFactory(faction string) AbstractFactory {
	switch faction {
	case "Elf":
		return ElfFactory{}
	case "Orc":
		return OrcFactory{}
	default:
		panic("Unknown faction")
	}
}

func main() {
	factory := getFactory("Elf")

	bow := factory.createBow()
	bow.attack()

	sword := factory.createSword()
	sword.attack()

	axe := factory.createAxe()
	axe.attack()
}
