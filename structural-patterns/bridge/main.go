package main

import "fmt"

/*
https://refactoring.guru/design-patterns/bridge

Bridge is a structural design pattern that lets you split a large class or a set of closely related classes
into two separate hierarchies—abstraction and implementation—which can be developed independently of each other.

Say you have a geometric Shape class with a pair of subclasses: Circle and Square.
You want to extend this class hierarchy to incorporate colors, so you plan to create Red and Blue
shape subclasses. However, since you already have two subclasses, you’ll need to create four class
combinations such as BlueCircle and RedSquare.

This problem occurs because we’re trying to extend the shape classes in two independent dimensions: by form and by color.
That’s a very common issue with class inheritance.

The Bridge pattern attempts to solve this problem by switching from inheritance to the object composition.

What this means is that you extract one of the dimensions into a separate class hierarchy,
so that the original classes will reference an object of the new hierarchy,
instead of having all of its state and behaviors within one class.

Following this approach, we can extract the color-related code into its own class with two subclasses: Red and Blue.
The Shape class then gets a reference field pointing to one of the color objects.
Now the shape can delegate any color-related work to the linked color object.

Abstraction (also called interface) is a high-level control layer for some entity.
This layer isn’t supposed to do any real work on its own.
It should delegate the work to the implementation layer (also called platform).

When talking about real applications, the abstraction can be represented by a graphical user interface (GUI),
and the implementation could be the underlying operating system code (API) which the GUI layer calls in response to user interactions.

Generally speaking, you can extend such an app in two independent directions:

1. Have several different GUIs (for instance, tailored for regular customers or admins).
2. Support several different APIs (for example, to be able to launch the app under Windows, Linux, and macOS).
In a worst-case scenario, this app might look like a giant spaghetti bowl, where hundreds of
conditionals connect different types of GUI with various APIs all over the code.

You can bring order to this chaos by extracting the code related to specific interface-platform combinations into separate classes.
However, soon you’ll discover that there are lots of these classes. The class hierarchy will grow exponentially because adding
a new GUI or supporting a different API would require creating more and more classes.

Let’s try to solve this issue with the Bridge pattern.
It suggests that we divide the classes into two hierarchies:

1. Abstraction: the GUI layer of the app.
2. Implementation: the operating systems’ APIs.

The abstraction object controls the appearance of the app, delegating the actual work to the linked implementation object.
Different implementations are interchangeable as long as they follow a common interface, enabling the same GUI to work under
Windows and Linux.

As a result, you can change the GUI classes without touching the API-related classes.
Moreover, adding support for another operating system only requires creating a subclass in the implementation hierarchy.
*/

/*
Say, you have two types of computers: Mac and Windows.
Also, two types of printers: Epson and HP.
Both computers and printers need to work with each other in any combination.
The client doesn’t want to worry about the details of connecting printers to computers.

If we introduce new printers, we don’t want our code to grow exponentially.
Instead of creating four structs for the 2*2 combination, we create two hierarchies:

1. Abstraction hierarchy: this will be our computers
2. Implementation hierarchy: this will be our printers

These two hierarchies communicate with each other via a Bridge,
where the Abstraction (computer) contains a reference to the Implementation (printer).
Both the abstraction and implementation can be developed independently without affecting each other.
*/

// Abstraction Class
type Computer interface {
	Print()
	SetPrinter(Printer)
}

// Refined Abstraction
type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

type Windows struct {
	printer Printer
}

func (m *Windows) Print() {
	fmt.Println("Print request for windows")
	m.printer.PrintFile()
}

func (m *Windows) SetPrinter(p Printer) {
	m.printer = p
}

// Implementation Class
type Printer interface {
	PrintFile()
}

// Refined Implementation
type Epson struct{}

func (e *Epson) PrintFile() {
	fmt.Println("Epson printer file")
}

type HP struct{}

func (h *HP) PrintFile() {
	fmt.Println("HP printer file")
}

func main() {
	mac := &Mac{}
	mac.SetPrinter(&HP{})
	mac.Print()
}
