
//interfaces
package main

import (
	"fmt"
)

type paymenter interface {
	pay(amount float32)
}

// payment struct uses the paymenter interface to abstract the payment gateway implementation.
// This is an example of Dependency Injection through interfaces.

type payment struct {
	gateway paymenter  
}



func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}



type jusPay struct{}

func (j jusPay) pay(amount float32) {
	fmt.Println("making payment using jusPay and amount is:- ", amount)
}



type fakePayment struct{}

func (f fakePayment) pay(amount float32) {
	fmt.Println("making payment using fakePayment for testing and amount is:- ", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32){
	fmt.Println("makeing payment using PayPay and amount is:- ",amount)
}


func main() {
	fakePaymentGW := fakePayment{}
	jusPayGW := jusPay{}
	paypalGW := paypal{}

	// Using fakePayment
	newPayment1 := payment{
		gateway: fakePaymentGW,
	}
	newPayment1.makePayment(100)

	// Using jusPay
	newPayment2 := payment{
		gateway: jusPayGW,
	}
	newPayment2.makePayment(200)

	// Using PayPal
	newPayment3 := payment{
		gateway: paypalGW,
	}
	newPayment3.makePayment(300)
}
