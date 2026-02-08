package main

import (
	"fmt"
	"time"
)

// struct are the similair to classes in other languages as they provide the way to take the multiple datatypes and put them in one. they are used to build complex datatrstructure, as the golang does not have the classes.
// example : order of an e-commrece app.

type order struct { // here the order is the new type which is created and can be used to create a new variable which has the same set of the values as this order type.
	id        string
	amount    float32
	status    string    // "paid" || "deleivered" || "cash"
	createdAt time.Time // A Time represents an instant in time with nanosecond precision.
}

// we can attach the methods as well to the struct as we have with the classes:
type order2 struct {
	id     string
	amount float32
	status string
}

func (o *order2) changeStatus(status string) { // (o order2) is the reciver type. the first letter of the struct is written as the name and the struct name is given before the function name in parenthesis.
	o.status = status // unused write to field status, this warning is given to handle that when the order2 struct type instance is passed it is passed by value which makes no change on the actaul instance when created. so we need to pass the order2 struct's instance by reference.
	// also we dont need to use the star * (dereference operator) here while accessing the status as it does not need to do that, struct does this internally.

}

func (o *order2) getAmount() float32 {
	return o.amount // we can also remove the reference while getting the value only. only need when we modify the value.
}

//-------------------------------------------------------------------
// golang does not have the constructor like the OOPS based language so we created a hack with which we wil create the method starting with 'new' keyword and then struct Name. this function will return the reference to the new struct instance created and returntype of the method is pointer of the struct name.

func newOrder(id string, amount float32, status string, createdAt time.Time) *order {
	order_init := order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: createdAt,
	}
	return &order_init // returning the address of the new struct instance created.
}

///////////////////--------main func----------------------

func main() {
	//--------------------- now we will create the instance of the struct we build.
	var order_one order = order{
		id:        "abc123",
		amount:    50.99,
		status:    "recieved",
		createdAt: time.Now(),
		// note : it is not compulsary to pass the all the values in the instance you can left some of these as well.
	}

	fmt.Println(order_one)

	// -------------------adding / replace the values later in the struct instance:
	// this is done using the dot operrator.
	order_one.id = "jkl1234"
	fmt.Println(order_one)

	// getting the value of the struct instance :
	fmt.Println(order_one.status)

	//------------------------- creating another struct: each instance is its own and does not effect other.
	order_two := order{
		id:     "apple123",
		amount: 45.66,
	}

	fmt.Println(order_two)

	//  --------------------------------------------- after the method is connected with struct

	order2_one := order2{
		id:     "word1323",
		amount: 9999.99,
		status: "cash",
	}
	fmt.Println(order2_one)
	// now we will simply call the method directly over the struct order2:
	order2_one.changeStatus("recieved")
	fmt.Println(order2_one)

	// another getter :
	fmt.Println(order2_one.getAmount())

	// --------------------------------------------

	// when we don't set any value for the value while creating the instance of the field then we will get the ZEROED value of the field. like 0 for the int, "" for string, 0 for the float etc.

	//-----------------------------------------------------

	// creating the new order instance with the method like constructor.
	order_three := newOrder("jamun123", 99.99, "recieved", time.Now())
	fmt.Println(order_three)

	//-------------------------------------
	// we can directly create the struct inside the main to just use it once when we have to store the configuration or not need to instace of it then.
	lang := struct {
		lang   string
		isGood bool
	}{"Golang", true} // same as struct in c lang.

	fmt.Println(lang)
}
