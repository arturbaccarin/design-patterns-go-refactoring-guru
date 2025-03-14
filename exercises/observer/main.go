package main

import (
	"fmt"
	"log"
)

/*
Create a system where we have a Subject (the thing being observed)
and multiple Observers that listen for changes in the Subject.
When the Subject changes, all registered Observers should be notified.
*/

type Observer interface {
	Update(message string)
}

type EmailObserver struct {
}

func (EmailObserver) Update(message string) {
	fmt.Printf("email update: %s", message)
}

type LoggingObserver struct {
}

func (LoggingObserver) Update(message string) {
	log.Printf("logging update: %s", message)
}

type Subject interface {
	Adding(o Observer)
	Removing(o Observer)
	Notify()
}

type NewMessage struct {
	observers []Observer
	message   string
}

func (n *NewMessage) Adding(o Observer) {
	n.observers = append(n.observers, o)
}

func (n *NewMessage) Removing(o Observer) {
	for i, observer := range n.observers {
		if observer == o {
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
		}
	}
}

func (n *NewMessage) Notify() {
	for _, observer := range n.observers {
		observer.Update(n.message)
	}
}

func main() {
	o1 := EmailObserver{}
	o2 := LoggingObserver{}

	newMessage := &NewMessage{
		observers: []Observer{o1, o2},
		message:   "New message",
	}

	newMessage.Notify()
}
