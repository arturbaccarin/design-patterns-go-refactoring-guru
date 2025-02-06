package main

import "fmt"

/*
You are building an online shopping system where the customer can browse products,
add them to the cart, apply discounts, and proceed to checkout. To make it easier
for customers to complete their orders, you want to simplify the interaction with
the complex subsystems involved in placing an order.

For this, we will implement the Facade design pattern, which will hide the
complexity of the subsystems (like inventory, discount, and payment) behind
a simpler interface.
*/

var products = map[int]int{
	1: 10,
	2: 0,
	3: 5,
	4: 0,
}

type ProductService struct{}

func (p *ProductService) CheckProductAvailaibility(id int) bool {
	quantity, ok := products[id]
	return ok && quantity > 0
}

func (p *ProductService) DecreaseProductQuantity(id int) {
	quantity, ok := products[id]
	if ok && quantity > 0 {
		products[id] = quantity - 1
	}
}

type DiscountService struct{}

func (d *DiscountService) ApplyDiscount(price float64) float64 {
	return price * 0.9
}

type PaymentService struct{}

func (p *PaymentService) ProcessPayment(amount float64) {
	fmt.Printf("Processing payment of %f\n", amount)
}

type OrderFacade struct {
	productService  *ProductService
	discountService *DiscountService
	paymentService  *PaymentService
}

func NewOrderFacade() *OrderFacade {
	return &OrderFacade{
		productService:  &ProductService{},
		discountService: &DiscountService{},
		paymentService:  &PaymentService{},
	}
}

func (o *OrderFacade) PlaceOrder(productId int, quantity int) {
	if !o.productService.CheckProductAvailaibility(productId) {
		fmt.Println("Product is not available")
	}

	o.productService.DecreaseProductQuantity(productId)
	price := float64(quantity) * 10
	discountedPrice := o.discountService.ApplyDiscount(price)
	o.paymentService.ProcessPayment(discountedPrice)
}

func main() {
	facade := NewOrderFacade()
	facade.PlaceOrder(1, 2)
}
