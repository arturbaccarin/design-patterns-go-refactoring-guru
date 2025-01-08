package main

import "fmt"

/*
https://refactoring.guru/design-patterns/memento

Memento is a behavioral design pattern that lets you save and restore the previous
state of an object without revealing the details of its implementation.

Imagine that you’re creating a text editor app. In addition to simple text editing, your editor can format text,
insert inline images, etc.

At some point, you decided to let users undo any operations carried out on the text.
This feature has become so common over the years that nowadays people expect every app
to have it. For the implementation, you chose to take the direct approach. Before performing
any operation, the app records the state of all objects and saves it in some storage. Later, when a
user decides to revert an action, the app fetches the latest snapshot from the history and uses
it to restore the state of all objects.

Let’s think about those state snapshots. How exactly would you produce one? You’d probably need
to go over all the fields in an object and copy their values into storage. However, this would
only work if the object had quite relaxed access restrictions to its contents. Unfortunately,
most real objects won’t let others peek inside them that easily, hiding all significant data in private fields.

In the future, you might decide to refactor some of the editor classes, or add or remove some of the fields.
Sounds easy, but this would also require changing the classes responsible for copying the state of the affected objects.

The Memento pattern delegates creating the state snapshots to the actual owner of that state, the originator object.
Hence, instead of other objects trying to copy the editor’s state from the “outside,” the editor class itself can
make the snapshot since it has full access to its own state.

The pattern suggests storing the copy of the object’s state in a special object called memento.
The contents of the memento aren’t accessible to any other object except the one that produced it.

Other objects must communicate with mementos using a limited interface which may allow fetching
the snapshot’s metadata (creation time, the name of the performed operation, etc.), but not
the original object’s state contained in the snapshot.

Such a restrictive policy lets you store mementos inside other objects, usually called caretakers.

In our text editor example, we can create a separate history class to act as the caretaker.
A stack of mementos stored inside the caretaker will grow each time the editor is about to execute an operation.

When a user triggers the undo, the history grabs the most recent memento from the stack
and passes it back to the editor, requesting a roll-back.
*/
type Originator struct {
	state string
}

func (e *Originator) createMemento() *Memento {
	return &Memento{state: e.state}
}

func (e *Originator) restoreMemento(m *Memento) {
	e.state = m.getSavedState()
}

func (e *Originator) setState(state string) {
	e.state = state
}

func (e *Originator) getState() string {
	return e.state
}

type Memento struct {
	state string
}

func (m *Memento) getSavedState() string {
	return m.state
}

type Caretaker struct {
	mementoArray []*Memento
}

func (c *Caretaker) addMemento(m *Memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *Caretaker) getMemento(index int) *Memento {
	return c.mementoArray[index]
}

func main() {

	caretaker := &Caretaker{
		mementoArray: make([]*Memento, 0),
	}

	originator := &Originator{
		state: "A",
	}

	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("B")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("C")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.restoreMemento(caretaker.getMemento(1))
	fmt.Printf("Restored to State: %s\n", originator.getState())

	originator.restoreMemento(caretaker.getMemento(0))
	fmt.Printf("Restored to State: %s\n", originator.getState())

}
