package main

import "fmt"

/*
https://refactoring.guru/design-patterns/state

State is a behavioral design pattern that lets an object alter its behavior
when its internal state changes. It appears as if the object changed its class.

The main idea is that, at any given moment, there’s a finite number of states which
a program can be in. Within any unique state, the program behaves differently,
and the program can be switched from one state to another instantaneously.

However, depending on a current state, the program may or may not switch
to certain other states. These switching rules, called transitions,
are also finite and predetermined.

You can also apply this approach to objects. Imagine that we have a Document class.
A document can be in one of three states: Draft, Moderation and Published.
The publish method of the document works a little bit differently in each state:

1. In Draft, it moves the document to moderation.
2. In Moderation, it makes the document public, but only if the current user is an administrator.
3. In Published, it doesn’t do anything at all.

State machines are usually implemented with lots of conditional statements (if or switch)
that select the appropriate behavior depending on the current state of the object.
Usually, this “state” is just a set of values of the object’s fields.

The biggest weakness of a state machine based on conditionals reveals itself
once we start adding more and more states and state-dependent behaviors to the Document class.
Most methods will contain monstrous conditionals that pick the proper behavior of a method
according to the current state. Code like this is very difficult to maintain because any
change to the transition logic may require changing state conditionals in every method.

The State pattern suggests that you create new classes for all possible states of an
object and extract all state-specific behaviors into these classes.

Instead of implementing all behaviors on its own, the original object, called context,
stores a reference to one of the state objects that represents its current state,
and delegates all the state-related work to that object.

To transition the context into another state, replace the active state object with
another object that represents that new state.

This is possible only if all state classes follow the same interface and the context
itself works with these objects through that interface.

This structure may look similar to the Strategy pattern, but there’s one key difference.
In the State pattern, the particular states may be aware of each other and initiate
transitions from one state to another, whereas strategies almost never know about each other.

Real-World Analogy
The buttons and switches in your smartphone behave differently depending on the current
state of the device:

1. When the phone is unlocked, pressing buttons leads to executing various functions.
2. When the phone is locked, pressing any button leads to the unlock screen.
3. When the phone’s charge is low, pressing any button shows the charging screen.

State is a behavioral design pattern that allows an object to change the behavior when its internal state changes.

Let’s apply the State Design Pattern in the context of vending machines.
That a vending machine can be in 4 different states:
1. hasItem
2. noItem
3. itemRequested
4. hasMoney

A vending machine will also have different actions. Again for simplicity lets assume that there are only four actions:
Select the item
1. Add the item
2. Insert money
3. Dispense item
*/

// Context
type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State

	currentState State

	itemCount int
	itemPrice int
}

func newVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &HasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}
	noItemState := &NoItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *VendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *VendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *VendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *VendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

func (v *VendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}

// State interface
type State interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

// Concrete state
type NoItemState struct {
	vendingMachine *VendingMachine
}

func (i *NoItemState) requestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (i *NoItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *NoItemState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}
func (i *NoItemState) dispenseItem() error {
	return fmt.Errorf("Item out of stock")
}

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (i *HasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}
	fmt.Printf("Item requestd\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *HasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *HasItemState) insertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}
func (i *HasItemState) dispenseItem() error {
	return fmt.Errorf("Please select item first")
}

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestedState) requestItem() error {
	return fmt.Errorf("Item already requested")
}

func (i *ItemRequestedState) addItem(count int) error {
	return fmt.Errorf("Item Dispense in progress")
}

func (i *ItemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("Inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}
func (i *ItemRequestedState) dispenseItem() error {
	return fmt.Errorf("Please insert money first")
}

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *HasMoneyState) requestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) addItem(count int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i *HasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}
func (i *HasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}
