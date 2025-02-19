package main

import "fmt"

type Light struct {
	isOn bool
}

func (l *Light) On() {
	l.isOn = true
	fmt.Println("Light is ON")
}

func (l *Light) Off() {
	l.isOn = false
	fmt.Println("Light is OFF")
}

type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}
