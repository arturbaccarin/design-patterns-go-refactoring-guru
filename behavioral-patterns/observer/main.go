package main

import "fmt"

/*
https://refactoring.guru/design-patterns/observer

Observer is a behavioral design pattern that lets you define a subscription
mechanism to notify multiple objects about any events that happen to the object
they’re observing.

Imagine that you have two types of objects: a Customer and a Store.
The customer is very interested in a particular brand of product.

The customer could visit the store every day and check product availability.
But while the product is still en route, most of these trips would be pointless.

On the other hand, the store could send tons of emails (which might be considered spam)
to all customers each time a new product becomes available.
This would save some customers from endless trips to the store.
At the same time, it’d upset other customers who aren’t interested in new products.

The object that has some interesting state is often called subject.
notify other objects about the changes to its state, we’ll call it publisher.

The Observer pattern suggests that you add a subscription mechanism to the publisher
class so individual objects can subscribe to or unsubscribe from a stream of events
coming from that publisher.

This mechanism consists of 1) an array field for storing a list of references to
subscriber objects and 2) several public methods which allow adding subscribers
to and removing them from that list.
*/

type EventType string

const (
	newProduct EventType = "newProduct"
	newOrder   EventType = "newOrder"
)

// The base publisher class includes subscription management
// code and notification methods.
type EventManager struct {
	listeners map[EventType][]Listener
}

func (e *EventManager) subscribe(eventType EventType, listener Listener) {
	if e.listeners == nil {
		e.listeners = make(map[EventType][]Listener)
	}
	e.listeners[eventType] = append(e.listeners[eventType], listener)
}

func (e *EventManager) unsubscribe(eventType EventType, listener Listener) {
	for i, l := range e.listeners[eventType] {
		if l == listener {
			e.listeners[eventType] = append(e.listeners[eventType][:i], e.listeners[eventType][i+1:]...)
		}
	}
}

func (e *EventManager) notify(eventType EventType, message string) {
	for _, listener := range e.listeners[eventType] {
		listener.update(message)
	}
}

type Listener interface {
	update(message string)
}

type Customer struct {
	name string
}

func (c *Customer) update(message string) {
	fmt.Printf("Customer %s received message: %s\n", c.name, message)
}

func main() {
	eventManager := &EventManager{}
	customer1 := &Customer{"Alice"}
	eventManager.subscribe(newProduct, customer1)
	eventManager.notify(newProduct, "New product available!")
	eventManager.unsubscribe(newProduct, customer1)
	eventManager.notify(newProduct, "New product available!")
}
