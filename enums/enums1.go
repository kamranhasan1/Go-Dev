

package main

import "fmt"

type orderStatus string//we can use int aslo 

const (
	received   orderStatus = "received"//insted or this we can write iota is will give the int value
	confirmed              = "confirmed"
	shipped                = "shipped"
	delivered              = "delivered"
)

func myOrderstatus(status orderStatus) {
	fmt.Println("your order has been", status)
}

func main() {
	myOrderstatus(delivered)
}