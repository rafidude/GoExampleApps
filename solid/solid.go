package main

// Single Responsibility Principle
// Payment Service
type PaymentService struct{}

func (p *PaymentService) ProcessPayment(amount float64) {
}

// Invoice Service
type InvoiceService struct{}

// CreateInvoice creates an invoice
func (i *InvoiceService) CreateInvoice(amount float64) {
}

// Open Closed Principle
// PaymentMethod is an interface
type PaymentMethod interface {
	Pay(amount float64) error
}

// CreditCard PaymentMethod is a struct
type CreditCard struct{}

// Pay is a method of CreditCard
func (c *CreditCard) Pay(amount float64) error {
	return nil
}

// PayPal PaymentMethod is a struct
type PayPal struct{}

// Pay is a method of PayPal
func (p *PayPal) Pay(amount float64) error {
	return nil
}

// ProcessPayment is a function
// func ProcessPayment(method PaymentMethod, amount float64) error {
// 	return method.Pay(amount)
// }

// Dependency Inversion Principle
// PayService is a struct
type PayService struct {
	method PaymentMethod
}

// ProcessPayment is a method of PayService
func (p *PayService) ProcessPayment(amount float64) error {
	return p.method.Pay(amount)
}

// SetPaymentMethod is a method of PayService
func (p *PayService) SetPaymentMethod(method PaymentMethod) {
	p.method = method
}
