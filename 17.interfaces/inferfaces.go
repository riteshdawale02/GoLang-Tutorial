package main

import "fmt"

type payment struct {
	gateway razorpay
}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGateway := razorpay{}
	// stripePaymentGateway := stripe{}
	p.gateway.pay(amount)
	// razorpayPaymentGateway.pay(amount)
}

type razorpay struct {}

func (r razorpay) pay(amount float32) {
	//logic to make payment
	fmt.Println("Making payment using razorpay:", amount)
}

type stripe struct {}

func (r stripe) payment(amount float32) {
	//logic to make payment
	fmt.Println("Making payment using stripe:", amount)
}

func main() {
	// stripePaymentGateway := stripe{}

	razorpayPaymentGateway := razorpay{}

	pay := payment{
		gateway: razorpayPaymentGateway,
	}
	pay.makePayment(200)


}
