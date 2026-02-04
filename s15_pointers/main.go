package main

import "fmt"

// pointers : it is the address in memory of the data structure stored in the memory.

// pass by value : num (a copy of num)
func changeNum(num int) {
	num = 5
	fmt.Println("In change num ", 5)
}

// pass by reference :
func changeNum2(num *int) {
	*num = 5 // de-reference
	fmt.Println("In change num ", 5)
}

func main() {
	var num int = 1
	changeNum(num)
	fmt.Println("after change the num in main ", num)
	// here the value of the num does not change after the function is being called. this is because, in golang, the values are pass by values not pass by reference.
	// to change them when passed, we will send the reference of the variable. ie. address of the variable.
	fmt.Println("address of the num", &num)
	changeNum2(&num)
	fmt.Println("after change the num in main ", num)
}
