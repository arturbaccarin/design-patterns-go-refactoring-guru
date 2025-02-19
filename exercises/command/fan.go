package main

import "fmt"

type Fan struct {
	isOn bool
}

func (f *Fan) On() {
	f.isOn = true
	fmt.Println("Fan is ON")
}

func (f *Fan) Off() {
	f.isOn = false
	fmt.Println("Fan is OFF")
}

type FanOnCommand struct {
	fan *Fan
}

func (c *FanOnCommand) Execute() {
	c.fan.On()
}

type FanOffCommand struct {
	fan *Fan
}

func (c *FanOffCommand) Execute() {
	c.fan.Off()
}
