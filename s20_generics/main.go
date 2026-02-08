package main

import "fmt"

//generics :

// ------- function with generics------------------

func printSlice(items []int) {
	for _, item := range items {
		fmt.Println("item : ", item)
	}
}

func printSlice2(items []string) {
	for _, item := range items {
		fmt.Println("item: ", item)
	}
}

// now, this is generics, where we can use any datatype of the slice to get it print.
func printSlice3[T any](items []T) {
	for _, item := range items {
		fmt.Println("item: ", item)
	}
}

// 'any' datatype for the T is not okay so we can set the multiple types using the pipe '|' which are allowed as:
func printSlice4[T string | int | bool](items []T) {
	for _, item := range items {
		fmt.Println("item: ", item)
	}
}

// in the above method we can only pass the string, int and bool.

// ------struct with generics-------------------
type stack[T any] struct { // here the stack is stack Data Structure which has approach of LIFO.
	elements []T
}

// this stack can used to store any data type we want.

// ----------another way of generics with function------------
//here we will use the word 'comparable' which is also an interface and contains many types.
// similary multiple generics can be used.
func printSlice5[T comparable, V string](items []T, value V) {
	for _, item := range items {
		fmt.Println("item: ", item)
	}
	fmt.Println(value)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	printSlice(nums)
	// now if we create the slice of string then we can't pass it in the printSlice method as it takes the type int slice.
	langs := []string{"typescript", "golang", "javascript"}
	// printSlice(langs) // can't do that.
	// so we need to create the another method which will print the string based slice.
	printSlice2(langs)
	// similarly for the boolean we need to create another method as well.
	// so, basically we are repeating the internal code loop of the function.
	// so, solve it we will use the generics.
	// (see above)

	printSlice3(nums)
	printSlice3(langs)

	//--------------generics with struct----------

	stack_one := stack[int]{ // we need to mention here what typeof the T is used when instance of the stack is created.
		elements: []int{1, 2, 3},
	}
	stack_two := stack[string]{
		elements: []string{"Apple", "Mango"},
	}
	fmt.Println(stack_one)
	fmt.Println(stack_two)

}
