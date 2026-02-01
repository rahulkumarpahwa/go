package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch:
	i := 3

	// we don't need to write break
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default: //optional
		fmt.Println("Other")
	}

	// multiple condition switch with multiple condition in case:
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's weekday")
	}

	//type switch using the switch within function
	whoAmI := func(i interface{}) { // empty interface shows that 'i' does not have any type. we can use the 'any' here as well instead of it.
		switch i.(type) { // i.(type) returns the type of the i
		case int:
			fmt.Println("It is an integer")
		case float32:
			fmt.Println("It is float")
		case string:
			fmt.Println("It is string")
		case bool:
			fmt.Println("It is bool")
		default:
			fmt.Println("It is other")

		}

	}

	whoAmI("apple")
	whoAmI(4)

}
