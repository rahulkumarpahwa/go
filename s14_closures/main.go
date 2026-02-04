package main

import "fmt"

func counter() func() int {
	var count int = 0

	return func() int { //this is increment fxn.
		count++
		return count
	}
	// the value of the normally variable inside the fxn gets deleted after the fxn declaration finishes but here it persist even after the outer fxn call in the main finishes to the inner fxn and increases everytime we call the inner method. this is possible due to closures.
}

func main() {

	increament_fxn := counter()
	fmt.Println(increament_fxn()) //1
	fmt.Println(increament_fxn()) //2
	fmt.Println(increament_fxn()) //3

	increament_fxn2 := counter()
	fmt.Println(increament_fxn2()) //1
	// although they refer to the same counter appear like that but actaully in memory a new counter value is created for the new fxn call to the Counter fxn itself which contains the fxn which increases the value.

}

// ============================
// ADDITIONAL NOTES ON CLOSURES
// ============================
//
// WHAT IS A CLOSURE?
// - A closure is a function that references variables from outside its body.
// - The function "closes over" these variables, meaning it can access and modify them.
//
// KEY CHARACTERISTICS:
// 1. Each closure has its own copy of the enclosed variables.
// 2. The enclosed variables persist in memory as long as the closure exists.
// 3. Normally, local variables are destroyed when a function returns.
//    But with closures, the variable persists because the returned function references it.
// 4. The inner function "captures" the variable, keeping it alive in memory.
//
// WHY CLOSURES ARE USEFUL:
// - Maintain private state without using global variables
// - Create function factories (functions that return customized functions)
// - Implement callbacks that remember their context
// - Build iterators and generators
//
// MEMORY NOTE:
// - increament_fxn and increament_fxn2 have SEPARATE 'count' variables.
// - Each call to counter() creates a new closure with its own enclosed state.
