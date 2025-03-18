package main

import "fmt"

/*
In this exercise, you'll create a payment system where different payment strategies (like CreditCard and PayPal) can be used to process payments.
*/

type PaymentStrategy interface {
	Pay(amount float64)
}

type CreditCardStrategy struct{}

func (c CreditCardStrategy) Pay(amount float64) {
	fmt.Printf("Paying %f using credit card\n", amount)
}

type PayPalStrategy struct{}

func (p PayPalStrategy) Pay(amount float64) {
	fmt.Printf("Paying %f using PayPal\n", amount)
}

type Context struct {
	strategy PaymentStrategy
}

func (c *Context) SetStrategy(strategy PaymentStrategy) {
	c.strategy = strategy
}

func (c *Context) DoSomething() {
	c.strategy.Pay(100.0)
}

func main() {
	context := Context{}
	context.SetStrategy(CreditCardStrategy{})
	context.DoSomething()
}
