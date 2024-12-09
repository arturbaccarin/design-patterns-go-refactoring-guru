package main

import "fmt"

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

With this new approach you can easily substitute the linked “helper” object with another, changing the behavior of the container at runtime.

“Wrapper” is the alternative nickname for the Decorator pattern that clearly expresses the main idea of the pattern.
A wrapper is an object that can be linked with some target object. The wrapper contains the same set of methods as the target and delegates
to it all requests it receives. However, the wrapper may alter the result by doing something either before or after it passes the request to the target.

When does a simple wrapper become the real decorator?
As I mentioned, the wrapper implements the same interface as the wrapped object.
That’s why from the client’s perspective these objects are identical.
*/

// Define the Operation interface
type Operation interface {
	Execute() error
}

// Concrete implementation of an Operation
type BaseOperation struct {
	Name string
}

func (b *BaseOperation) Execute() error {
	// Simulate some work
	fmt.Println("Executing operation:", b.Name)
	return nil
}

// Decorator that adds behavior before and after the original execution
type OperationDecorator struct {
	Operation Operation // The wrapped operation
}

func (d *OperationDecorator) Execute() error {
	// Pre-processing: logging before execution
	fmt.Println("Before execution:", d.Operation)

	// Execute the original operation
	err := d.Operation.Execute()

	// Post-processing: logging after execution
	fmt.Println("After execution:", d.Operation)

	return err
}

func main() {
	// Create a concrete operation
	operation := &BaseOperation{Name: "My Operation"}

	// Wrap the operation with a decorator
	decoratedOperation := &OperationDecorator{Operation: operation}

	// Execute the decorated operation
	decoratedOperation.Execute()
}
