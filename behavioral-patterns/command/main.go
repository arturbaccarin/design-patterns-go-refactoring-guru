package main

import "fmt"

/*
https://refactoring.guru/design-patterns/command

Command is a behavioral design pattern that turns a request into a stand-alone
object that contains all information about the request.

This transformation lets you pass requests as a method arguments, delay or
queue a request’s execution, and support undoable operations.

Good software design is often based on the principle of separation of concerns,
which usually results in breaking an app into layers.
The most common example: a layer for the graphical user interface and another layer for
the business logic. The GUI layer is responsible for rendering a beautiful
picture on the screen, capturing any input and showing results of what the user and the app are doing.

However, when it comes to doing something important, like calculating the trajectory
of the moon or composing an annual report, the GUI layer delegates the work to
the underlying layer of business logic.

The Command pattern suggests that GUI objects shouldn’t send these requests directly.

Instead, you should extract all of the request details, such as the object being called,
the name of the method and the list of arguments into a separate command class
with a single method that triggers this request.

Command objects serve as links between various GUI and business logic objects.

From now on, the GUI object doesn’t need to know what business logic object will
receive the request and how it’ll be processed. The GUI object just triggers the command,
which handles all the details.

The next step is to make your commands implement the same interface.
Usually it has just a single execution method that takes no parameters.

Since the command execution method doesn’t have any parameters, how would we pass the
request details to the receiver? It turns out the command should be either
pre-configured with this data, or capable of getting it on its own.

They’ll be linked to a command which gets executed when a user interacts with the GUI element.
As you’ve probably guessed by now, the elements related to the same operations will
be linked to the same commands, preventing any code duplication.

As a result, commands become a convenient middle layer that reduces coupling
between the GUI and business logic layers.

After a long walk through the city, you get to a nice restaurant and sit at
the table by the window. A friendly waiter approaches you and quickly takes your order,
writing it down on a piece of paper.
The waiter goes to the kitchen and sticks the order on the wall.
After a while, the order gets to the chef, who reads it and cooks the meal accordingly.
The cook places the meal on a tray along with the order.
The waiter discovers the tray, checks the order to make sure everything is as you wanted it,
and brings everything to your table.

The paper order serves as a command. It remains in a queue until the chef is ready to
serve it. The order contains all the relevant information required to cook the meal.
It allows the chef to start cooking right away instead of running around clarifying
the order details from you directly.

I`m going back tomorrow! Thank you a happy new year!

Command is behavioral design pattern that converts requests or simple operations into objects.

Let’s look at the Command pattern with the case of a TV. A TV can be turned ON by either:
    1. ON Button on the remote;
    2. ON Button on the actual TV.

We can start by implementing the ON command object with the TV as a receiver.

The last part is defining an invoker. We’ll actually have two invokers: the remote and the TV itself.

There’s no need to develop different handlers for each of the invokers.
The command object contains all the information it needs to execute.
Hence it can also be used for delayed execution.
*/

// Invoker
type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

type Command interface {
	execute()
}

// Concrete command
type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

// Receiver interface
type Device interface {
	on()
	off()
}

// Concrete receiver
type TV struct {
	isRunning bool
}

func (t *TV) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *TV) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func main() {
	tv := &TV{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}
