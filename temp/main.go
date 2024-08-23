package main

import "fmt"

type Paymenter interface {
	pay()
}

type Payment struct {
	gateway Paymenter
}

type Razorpay struct {
	amount float32
}

func (r Razorpay) pay() {
	fmt.Println("Doing payment through Razorpay")
}

type Stripe struct {
	amount float32
}

func (s Stripe) pay() {
	fmt.Println("Doing payment through Stripe")
}

func main() {
	//
	razorPayGW := Razorpay{amount: 202}
	stripePayGW := Stripe{amount: 202}
	PaymentOne := Payment{gateway: razorPayGW}
	PaymentTwo := Payment{gateway: stripePayGW}
	PaymentOne.gateway.pay()
	PaymentTwo.gateway.pay()
}
