package main

/*
In this exercise, you'll create a simple smart home system that allows users
to control various appliances like lights, fans, and the thermostat using the
Command Pattern. You will define commands for turning the devices on or off,
and implement an invoker that handles the requests.
*/

type Command interface {
	Execute()
}

type RemoteControl struct {
	commands []Command
}

func (r *RemoteControl) SetCommand(c Command) {
	r.commands = append(r.commands, c)
}

func (r *RemoteControl) PressButton() {
	if len(r.commands) > 0 {
		command := r.commands[len(r.commands)-1]
		r.commands = r.commands[:len(r.commands)-1]
		command.Execute()
	}
}

func main() {
	// Create appliances
	light := &Light{}
	fan := &Fan{}
	thermostat := &Thermostat{}

	// Create commands
	lightOn := &LightOnCommand{light}
	lightOff := &LightOffCommand{light}
	fanOn := &FanOnCommand{fan}
	fanOff := &FanOffCommand{fan}
	setThermostat := &ThermostatSetCommand{thermostat, 22}

	// Create remote control (Invoker)
	remote := &RemoteControl{}

	// Use remote to control appliances
	remote.SetCommand(lightOn)
	remote.PressButton()

	remote.SetCommand(fanOn)
	remote.PressButton()

	remote.SetCommand(setThermostat)
	remote.PressButton()

	remote.SetCommand(lightOff)
	remote.PressButton()

	remote.SetCommand(fanOff)
	remote.PressButton()
}
