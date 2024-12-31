package main

import (
	"fmt"
	"time"
)

type customer struct{
	name string
	phone int
}
//order struct
type order struct{
	id string
	amount float32
	status string
	createdAt time.Time // nanoseconf precision
	customer // struct embedding
}

func newOrder(id string, amount float32, status string) *order{
	myOrder := order{
		id: id,
		amount:amount ,
		status: status,
	}

	return &myOrder
}
// 0 order now we have connected to order struct
func (o *order) changestatus(status string){
	o.status = status // this wouldnt change the instance value
	// we neet=d to add *  before order which means pointer 
}

func main() {
	// var order order = 
	// myOrder := order{
	// 	id: "1",
	// 	amount: 50.00,
	// 	status: "received",
	// }

	newCustomer := customer{
		name: "new",
		phone: 2323232323,
	}

	myOrder := order{
		id: "1",
		amount: 50.00,
		status: "received",
		customer: newCustomer,
	}

	// myOrder.createdAt = time.Now()
	// fmt.Println(myOrder.id)
	
	// myOrder.changestatus("Confirmed")
	// fmt.Println("order struct", myOrder)
	fmt.Println("order struct", myOrder) 

	// o/p: order struct {1 50 Confirmed {0 0 <nil>}}

	// myOrder := newOrder("1", 50.50, "confiremd")
	// fmt.Println(myOrder)

	// inline struct

	language := struct{
		name  string
		isbool bool
	}{"GO", true}
	fmt.Println(language)
	
	
}