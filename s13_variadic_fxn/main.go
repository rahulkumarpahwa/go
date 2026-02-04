package main

import "fmt"

// when we can pass the n number of parameters in the fxn then it is called variadic fxns.
// same as spread and rest operator in js.
func sum(nums ...int) int { // here the nums act as the slice.
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}

// to create the same fxn which can take any type, then we can either use the interface{} or the any type.
func anyViewer(anyValues ...interface{}) any { // or use any
	for _, val := range anyValues {
		fmt.Println(val)
	}
	return nil
}

func main() {
	fmt.Println(sum(4, 5, 5, 6, 6, 6, 6, 6))
	fmt.Println(anyViewer("apple", "mango", "golang", "js"))

	// passing slice in the variadic fxn:
	nums := []int{3, 4, 5, 6}
	fmt.Println(sum(nums...)) // spread as js but opposite.

}
