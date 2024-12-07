package main

/*
https://refactoring.guru/design-patterns/decorator

Decorator is a structural design pattern that lets you attach new behaviors to objects
by placing these objects inside special wrapper objects that contain the behaviors.

Extending a class is the first thing that comes to mind when you need to alter an object’s behavior.
However, inheritance has several serious caveats that you need to be aware of.

1. 	Inheritance is static. You can’t alter the behavior of an existing object at runtime.
	You can only replace the whole object with another one that’s created from a different subclass.

2.	Subclasses can have just one parent class. In most languages,
	inheritance doesn’t let a class inherit behaviors of multiple classes at the same time.

One of the ways to overcome these caveats is by using Aggregation or Composition instead of Inheritance.
Both of the alternatives work almost the same way: one object has a reference to another and delegates it some work,
whereas with inheritance, the object itself is able to do that work, inheriting the behavior from its superclass.
*/
