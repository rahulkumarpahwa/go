package main

import "fmt"

const name2 string = "constant value"

// name3 := "not allowed this way outside fxn"
var name3 string = "this is also allowed"

func main() {
	const name string = "apple"
	fmt.Println("constant value is " + name)
	// name = "javascript"; // not allowed, constant.
	fmt.Println(name2)
	fmt.Println(name3)

	// create the multiple constant as in one group using:
	const (
		port int    = 5000
		host string = "localhost"
	)

	fmt.Println(port, host)
}
