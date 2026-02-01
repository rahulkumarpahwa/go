package main

import "fmt"

//numbered sequence of specific length ie. length will be fixed and will be of same type.

func main() {
	var nums [10]int
	fmt.Println(len(nums)) //length
	nums[0] = 10           // setting 10 at index 0
	fmt.Println(nums[0])   //get the value at index 0
	fmt.Println(nums)      // full array with zeroed values

	// zeroed values are the ones which are putted in the array inplace of the empty value. eg. zero of the int, empty string of string array, false for bool array, etc.

	// to declare in the single line.
	var nums2 [3]int = [3]int{1, 2, 3}
	nums4 := [3]int{1, 2, 3} // other way
	fmt.Println(nums2)
	fmt.Println(nums4)

	//2d array:
	var nums3 [2][2]int = [2][2]int{{3, 4}, {5, 6}}
	fmt.Println(nums3)

	// notes :
	// - mostly slices are used as comapared to array due to fixed size before declaration.
	// -  fixed size, that is predictable
	// - memory optimisation
	// - constant time access
}
