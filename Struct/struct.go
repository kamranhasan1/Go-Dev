package main

import (
	"fmt"
)

type Order struct {
	product     string
	price       int
	trackingId  int
	status      string
}

func (o *Order) changeStatus(status string){ //reciver type
	o.status=status

}


func main() {
	myOrder := Order{
		product:    "wiper",
		price:      23,
		trackingId: 1224,
		status:     "shipped",
	}
	fmt.Println(myOrder)
	myOrder.changeStatus("dispached")
	fmt.Println(myOrder)
}
