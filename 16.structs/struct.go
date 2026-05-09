package main

import (
	"fmt"
	"time"
)

//order struct
 type order struct {
	id string
	amount float32
	status string
	createdAt time.Time  //nanosecond precision
}

//Struct Embedding -> adding struct in struct nested struct
 type orderStruct struct {
	id string
	amount float32
	status string
	createdAt time.Time  //nanosecond precision
	customer
}

type customer struct {
	name string
	number string
}


func newOrder(id string, amount float32, status string, createdAt time.Time) *order {

	//initial set goes here...
	myOrder := order {
		id: id,
		amount: amount,
		status: status,
		createdAt: createdAt,
	}

	return &myOrder
}


func (o *order) changeStatus(status string) {
	o.status = status
}

func main() {
	myOrder := order {
		id: "1",
		amount: 50.33,
		status: "Paid",
	}

	myOrder1 := newOrder("4",12.60, "received", time.Now())
	fmt.Println(myOrder1.amount)
	fmt.Println("================")


	myOrder.changeStatus("Failed")

	myOrder2 := order {
		id: "2",
		amount: 100.23,
		status: "Delivered",
		createdAt: time.Now(),
	}

	myOrder.createdAt = time.Now()
	fmt.Println(myOrder.amount)
	fmt.Println(myOrder)
	fmt.Println(myOrder2)



	language := struct {
		name string
		isGood bool
	} {"goLang", true}

	fmt.Println(language)



	nestedStruct:= orderStruct{
		id: "6",
		amount: 20.40,
		status: "failed",
		createdAt: time.Now(),
		customer: customer{
			name: "ritesh",
			number: "9699838302",
		},
	}
	nestedStruct.customer.name = "Samruddhi"

	fmt.Println(nestedStruct)
}