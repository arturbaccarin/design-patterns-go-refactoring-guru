package main

import "fmt"

/*
https://refactoring.guru/design-patterns/mediator

Mediator is a behavioral design pattern that lets you reduce chaotic
dependencies between objects.

The pattern restricts direct communications between the objects and
forces them to collaborate only via a mediator object.

Say you have a dialog for creating and editing customer profiles.
It consists of various form controls such as text fields, checkboxes, buttons, etc.

Some of the form elements may interact with others. For instance, selecting the
“I have a dog” checkbox may reveal a hidden text field for entering the dog’s name.
Another example is the submit button that has to validate values of all fields
before saving the data.

By having this logic implemented directly inside the code of the form elements you make
these elements’ classes much harder to reuse in other forms of the app.

For example, you won’t be able to use that checkbox class inside
another form, because it’s coupled to the dog’s text field.

The Mediator pattern suggests that you should cease all direct communication between the
components which you want to make independent of each other.

Instead, these components must collaborate indirectly, by calling a special mediator
object that redirects the calls to appropriate components.
As a result, the components depend only on a single mediator class instead of being
coupled to dozens of their colleagues.

The most significant change happens to the actual form elements.
Let’s consider the submit button. Previously, each time a user clicked
the button, it had to validate the values of all individual form elements.
Now its single job is to notify the dialog about the click.
Upon receiving this notification, the dialog itself performs the validations
or passes the task to the individual elements. Thus, instead of being tied
to a dozen form elements, the button is only dependent on the dialog class.

Pilots of aircraft that approach or depart the airport control area
don’t communicate directly with each other. Instead, they speak to
an air traffic controller, who sits in a tall tower somewhere near
the airstrip. Without the air traffic controller, pilots would need
to be aware of every plane in the vicinity of the airport, discussing
landing priorities with a committee of dozens of other pilots.

The tower doesn’t need to control the whole flight. It exists only
to enforce constraints in the terminal area because the number of
involved actors there might be overwhelming to a pilot.
*/
// Mediator Interface
type Mediator interface {
	Notify(sender Colleague, event string)
}

// Colleague Interface
type Colleague interface {
	SetMediator(mediator Mediator)
	Send(event string)
	Receive(event string)
}

// Concrete Mediator
type ConcreteMediator struct {
	colleague1 Colleague
	colleague2 Colleague
}

func (m *ConcreteMediator) SetColleague1(c Colleague) {
	m.colleague1 = c
}

func (m *ConcreteMediator) SetColleague2(c Colleague) {
	m.colleague2 = c
}

func (m *ConcreteMediator) Notify(sender Colleague, event string) {
	if sender == m.colleague1 {
		m.colleague2.Receive(event)
	} else {
		m.colleague1.Receive(event)
	}
}

// Concrete Colleague 1
type ConcreteColleague1 struct {
	mediator Mediator
}

func (c *ConcreteColleague1) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *ConcreteColleague1) Send(event string) {
	fmt.Println("Colleague 1 sends:", event)
	c.mediator.Notify(c, event)
}

func (c *ConcreteColleague1) Receive(event string) {
	fmt.Println("Colleague 1 receives:", event)
}

// Concrete Colleague 2
type ConcreteColleague2 struct {
	mediator Mediator
}

func (c *ConcreteColleague2) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *ConcreteColleague2) Send(event string) {
	fmt.Println("Colleague 2 sends:", event)
	c.mediator.Notify(c, event)
}

func (c *ConcreteColleague2) Receive(event string) {
	fmt.Println("Colleague 2 receives:", event)
}

func main() {
	mediator := &ConcreteMediator{}

	colleague1 := &ConcreteColleague1{}
	colleague2 := &ConcreteColleague2{}

	// Set the mediator for both colleagues
	colleague1.SetMediator(mediator)
	colleague2.SetMediator(mediator)

	// Set up the relationship in the mediator
	mediator.SetColleague1(colleague1)
	mediator.SetColleague2(colleague2)

	// Colleague 1 sends a message
	colleague1.Send("Hello from Colleague 1")

	// Colleague 2 sends a message
	colleague2.Send("Hello from Colleague 2")
}
