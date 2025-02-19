package main

import "fmt"

type Thermostat struct {
	temperature int
}

func (t *Thermostat) SetTemperature(temp int) {
	t.temperature = temp
	fmt.Printf("Thermostat set to %d degrees\n", temp)
}

type ThermostatSetCommand struct {
	thermostat *Thermostat
	temp       int
}

func (c *ThermostatSetCommand) Execute() {
	c.thermostat.SetTemperature(c.temp)
}
