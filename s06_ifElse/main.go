package main

import "fmt"

func main() {
	// if else
	age := 18
	if age >= 18 {
		fmt.Println("Person is an adult ", age)
	} else {
		println("Person is not adult ", age)
	}

	// else if
	if age >= 18 {
		fmt.Println("Person is an adult ", age)
	} else if age < 18 {
		println("Person is not adult ", age)
	} else {
		println("Person is a kid")
	}

	var role string = "admin"
	var hasPermissions bool = false

	// logical operators:
	if role == "admin" || hasPermissions {
		fmt.Println("Can Access!")
	}

	if role == "dmin" && hasPermissions {
		fmt.Println("can Access!")
	}

	// declration of variable in the if directly.
	if age := 15; age >= 18 {
		fmt.Println("persom is an adult", age)
	} else if age >= 12 {
		fmt.Println("person is teenager", age)
	}

	// go does not have a ternary operator you have to use the normal if else.
}
