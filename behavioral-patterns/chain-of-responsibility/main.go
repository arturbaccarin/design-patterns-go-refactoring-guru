package main

import "fmt"

/*
https://refactoring.guru/design-patterns/chain-of-responsibility

Chain of Responsibility is a behavioral design pattern that lets you pass requests along a chain of handlers.
Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.

Imagine that you’re working on an online ordering system.
You want to restrict access to the system so only authenticated users can create orders.
Also, users who have administrative permissions must have full access to all orders.

After a bit of planning, you realized that these checks must be performed sequentially.
The application can attempt to authenticate a user to the system whenever it receives a request that contains the user’s credentials.
However, if those credentials aren’t correct and authentication fails, there’s no reason to proceed with any other checks.

During the next few months, you implemented several more of those sequential checks.

* One of your colleagues suggested that it’s unsafe to pass raw data straight to the ordering system.
So you added an extra validation step to sanitize the data in a request.

* Later, somebody noticed that the system is vulnerable to brute force password cracking.
To negate this, you promptly added a check that filters repeated failed requests coming from the same IP address.

* Someone else suggested that you could speed up the system by returning cached results on repeated requests containing
the same data. Hence, you added another check which lets the request pass through to the system only if
there’s no suitable cached response.

The code of the checks, which had already looked like a mess, became more and more bloated as you added each new feature.
Changing one check sometimes affected the others.

Worst of all, when you tried to reuse the checks to protect other components of the system,
you had to duplicate some of the code since those components required some of the checks, but not all of them.

Like many other behavioral design patterns, the Chain of Responsibility relies on
transforming particular behaviors into stand-alone objects called handlers.
In our case, each check should be extracted to its own class with a single method that performs the check.
The request, along with its data, is passed to this method as an argument.

The pattern suggests that you link these handlers into a chain.
Each linked handler has a field for storing a reference to the next handler in the chain.

a handler can decide not to pass the request further down the chain and effectively stop any further processing.

However, there’s a slightly different approach (and it’s a bit more canonical) in which, upon receiving a request,
a handler decides whether it can process it. If it can, it doesn’t pass the request any further.
So it’s either only one handler that processes the request or none at all.
This approach is very common when dealing with events in stacks of elements within a graphical user interface.

For instance, when a user clicks a button, the event propagates through the chain of GUI elements that
starts with the button, goes along its containers (like forms or panels),
and ends up with the main application window.

It’s crucial that all handler classes implement the same interface.
*/

/*
A hospital app. A hospital could have multiple departments such as:

1. Reception
2. Doctor
3. Medical examination room
4. Cashier

The patient is being sent through a chain of departments, where each department sends
the patient further down the chain once their function is completed.

The pattern is applicable when there are multiple candidates to process the same request.
It is also useful when you don’t want the client to choose the receiver as there are multiple objects can handle the request.
Another useful case is when you want to decouple the client from receivers—the client will only need to know the first element in the chain.
*/

type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

// Handler interface
type Department interface {
	execute(*Patient)
	setNext(Department)
}

// Concrete handler
type Reception struct {
	next Department
}

func (r *Reception) setNext(next Department) {
	r.next = next
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

type Doctor struct {
	next Department
}

func (d *Doctor) setNext(next Department) {
	d.next = next
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

type Medical struct {
	next Department
}

func (m *Medical) setNext(next Department) {
	m.next = next
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

type Cashier struct {
	next Department
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func main() {

	cashier := &Cashier{}

	//Set next for medical department
	medical := &Medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &Doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)
}
