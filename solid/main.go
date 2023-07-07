package main

import "fmt"

// Good examples

type PaymentMethod interface {
	Pay(amount float64)
}

type CreditCard struct{}

func (cc CreditCard) Pay(amount float64) {
	// Process credit card payment
	fmt.Println("CreditCard")
}

type PayPal struct{}

func (pp PayPal) Pay(amount float64) {
	// Process PayPal payment
	fmt.Println("PayPal")
}

type PaymentService struct {
	method PaymentMethod
}

func (ps *PaymentService) SetPaymentMethod(method PaymentMethod) {
	ps.method = method
}

func (ps PaymentService) ProcessPayment(amount float64) {
	ps.method.Pay(amount)
}

func main() {
	cc := CreditCard{}
	pp := PayPal{}

	paymentService := PaymentService{}

	paymentService.SetPaymentMethod(cc)
	paymentService.ProcessPayment(100)

	paymentService.SetPaymentMethod(pp)
	paymentService.ProcessPayment(200)
}
