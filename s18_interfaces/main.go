package main

import "fmt"

// Interfaces : used in oops mainly and in go as well.
// we will intregate the user payment gateway in the app.

type payment struct {
}

func (p payment) makePayment(amount float32) {
	// razorPayPaymentGateWay := razorPay{}
	// razorPayPaymentGateWay.pay(amount)

	stripePaymentGateway := stripe{}
	stripePaymentGateway.pay(amount)
}

type razorPay struct {
}

func (r razorPay) pay(amount float32) {
	//logic to make payment
	fmt.Println("making payment through the razorpay with amount : ", amount)
}

// till now this was simple to add a payment gateway which is razorpay here. but in future if someone wants the other payment gateway say "stripe" then we will make the struct of the stripe and then create the same method make with stripe struct of pay.

type stripe struct {
}

func (s stripe) pay(amount float32) {
	//logic to make payment
	fmt.Println("making payment through the stripe with amount : ", amount)
}

// but this would lead us to change the makePayment of the payment struct every time a new payment gateway comes and creates the tight coupling.
// we are voilating the principle of open-close : which says that classes should be open extension but closed for modification.

// so we will write this :
type payment2 struct {
	gateway stripe
	// we are here creating the struct of the stripe inside the struct of payment, this is what composition is.
}

// and now we can access this as :
func (p payment2) makePayment2(amount float32) {
	p.gateway.pay(amount)
}

// -------------------------interfaces---------
// interface names usally ends with "er" as it is normal convention.
// this is normally as the interfaces in the java.
// this is like the contract that every struct who follow this interface will have to have this method.
type paymenter interface {
	pay(amount float32) // return type can also be mention here.
}

// now we will create the payment struct but with the paymenter type so that any payment gateway can be used with as :
type payment3 struct {
	gateway paymenter
}

func (p payment3) makePayment3(amount float32) {
	p.gateway.pay(amount)
}

func main() {

	new_payment := payment{}
	new_payment.makePayment(100)

	// now we need to create the instance of the stripe and pass it in the payment2 struct instance as this what we have done previous to create the variable of type stripe struct in the payment.
	stripe_one := stripe{}

	new_payment2 := payment2{
		gateway: stripe_one,
	}
	new_payment2.makePayment2(101)

	// still when a new payment gateway needs to be added then we need to modify the payment2 struct to have the gateway of variable stripe or any other whose gateway we want then we will create the instance of that gateway and pass in the payment and then call the method of makePayment over the struct. so still our problem persist.
	// so to solve this we have interfaces. (See top)

	// now we will when create the stripe or razorpay or anyother payment gateway and pass it without thinking in the payment3 struct instance as it is not tightly bound.

	stripe_two := stripe{}
	new_payment3 := payment3{gateway: stripe_two}
	new_payment3.makePayment3(100)
	razorPay_two := razorPay{}
	new_payment4 := payment3{razorPay_two}
	new_payment4.makePayment3(3000)


	//note : in golang, we don't need to explicitly define the interface with "implements" keyword. go-compiler can directly understand that if struct as the same 'signature' method as the one defined in the inteface then it struct is implementing the interface.
	// now this works with open close principle.
	// here we have done the dependency inversion as we are not depending on any gateway and any type of gateway can be passed here without changing the payment struct and just need to add the new methods in the interface so that any gateway that comes follow it. 
	// these concept makes our app scalable. 
}
