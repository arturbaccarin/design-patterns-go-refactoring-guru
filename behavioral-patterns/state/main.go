package main

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
