package main

import "fmt"

// we must have the return type mention as well. mostly like the typescriptexcept colon.
func add(a int, b int) int {
	return a + b

}

// we can define the combined datatype of the parameter of the same type at once by giving the return type in end.
func add2(a, b int) int {
	return a + b
}

// we can return the multiple value from the fxn in go:
func get_languages() (string, string, string, bool) {
	return "javascript", "golang", "java", true
}

// functions are the first class citizen in go. so we can pass the function in variable or the fxn in other fxn.
// passing the fxn in fxn:
func InputTakerFxn(inputFxn func(a int) int) {
	fmt.Println(inputFxn(3)) // this fxn takes the parameter as int so we need to pass here and return the int.
}

// fxn returning the fxn:
func OutputGiverFxn() func(a int) int {
	return func(a int) int {
		return a * 4
	}
}

func main() {
	result := add(4, 5)
	fmt.Println(result)

	result2 := add2(5, 6)
	fmt.Println(result2)

	//  getting multiple return values from the fxn.
	fmt.Println(get_languages())
	lang1, lang2, lang3, _ := get_languages() // fourth value ignore by _
	fmt.Println(lang1, lang2, lang3)

	// calling the taker fxn:
	in_fxn := func(a int) int { // anonymous fxn
		return a * 2
	}

	InputTakerFxn(in_fxn) // above the fxn is defined and then passed.

	out_fxn := OutputGiverFxn() // this fxn gives the fxn out which can be called with paramter it needs
	fmt.Println(out_fxn(45))

}
