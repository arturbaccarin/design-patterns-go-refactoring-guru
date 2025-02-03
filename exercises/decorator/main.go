package main

import (
	"fmt"
	"time"
)

/*
In this exercise, you will implement a decorator in Go that measures
the execution time of any given function. The decorator will wrap around
the function and output the time it took for the function to execute.
*/

type Operation interface {
	Execute() error
}

type AnyFunction interface {
	Execute() error
}

type MyFunction struct {
	Name string
}

func (f *MyFunction) Execute() error {
	time.Sleep(1 * time.Second)
	fmt.Println("Executing operation:", f.Name)
	return nil
}

type OperationDecorator struct {
	Operation AnyFunction
}

func (d *OperationDecorator) Execute() error {
	start := time.Now()
	err := d.Operation.Execute()
	duration := time.Since(start)
	fmt.Printf("Operation %s took %s to execute\n", d.Operation, duration)
	return err
}

func Execute(operation Operation) error {
	operation.Execute()
	return nil
}

func main() {
	operation := &MyFunction{Name: "My Operation"}

	decoratedOperation := &OperationDecorator{Operation: operation}

	// Execute(operation)
	Execute(decoratedOperation)
}
