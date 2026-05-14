package main

import "fmt"

//enumrated types

const()

//we can define it as string
type OrderStatus int 

const (
	Received OrderStatus = iota
	Confirmed  // = "Confirmed"
	Prepared   // = "Prepared"
	Delivered  // = "Delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func main() {
	fmt.Println("Hello")
	changeOrderStatus(Received)
}