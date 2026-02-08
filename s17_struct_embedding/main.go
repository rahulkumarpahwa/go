package main

import (
	"fmt"
	"time"
)

// in oops, we have the concept of the inheritance and composition. same in oops we have the struct embeddings. eg. we can have the one struct inside the other ie. customer struct inside the order struct store the customer details in order struct.

type customer struct {
	name  string
	phone string
}

// now this function is also accessible in the order struct as the customer struct is embedded.
func (c *customer) changePhone(s string) {
	c.phone = s
}

type order struct {
	id     string
	amount float32
	status string
	time   time.Time
	cust   customer // this is called as the struct embedding or the composition as per oops.
}

func main() {
	order_one := order{
		id:     "apple1",
		amount: 344,
		status: "recived",
	}

	fmt.Println(order_one) // o/p : {apple1 344 recived {0 0 <nil>} { }}
	// you will see a empty {} at the end which shows the two empty strings for the customer details.

	// we can set the values of the embedded using the dot operator as well:
	order_one.cust.name = "apple"
	order_one.cust.phone = "987733737373"
	fmt.Println(order_one)

	// we can pre create the customer and then pass it as well.
	cust_two := customer{name: "mango", phone: "999999999"}
	order_two := order{
		id:   "7373737",
		cust: cust_two,
	}
	fmt.Println(order_two)

	// we can do this as well :
	order_three := order{
		status: "checkout",
		cust: customer{
			"mango", "848484848",
		},
	}
	fmt.Println(order_three)

	// we can use the embedded struct method in the struct where it is embedded.
	order_three.cust.changePhone("0000000000")
	fmt.Println(order_three)
}
