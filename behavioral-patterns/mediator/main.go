package main

/*
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
type Mediator interface {
	notify(sender Component, event string)
}

type Component interface {
	setMediator(mediator Mediator)
}

type Button struct {
	mediator Mediator
}

type Textbox struct {
	mediator Mediator
}

type Dialog struct {
	title string
	loginOrRegisterChkBx
}
