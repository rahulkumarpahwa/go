package main

import "fmt"

// for : only construct for the looping in Golang
func main() {
	// while loop : written using the for-loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// find diff between Println and println
	// infinite loop
	// for {
	// 	println("1")
	// }

	// classic for loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// we can do that using the range as well:
	for i := range 3 {
		// note : last value is not printed with range.
		fmt.Println(i)
	}

	var count int = 9
	// break and continue:
	for i := 0; i < count; i++ {
		if i == 5 {
			break
		}
		fmt.Println("value of i : ", i)
	}

	for i := 0; i < count; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("value of i : ", i)
	}

}
