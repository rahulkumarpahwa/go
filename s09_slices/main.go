package main

import (
	"fmt"
	"slices"
)

/* slices : 1. dynamic array
2. most useful construct in go
3. useful methods like arrays
*/

func main() {

	var nums []int           // unintialiosed slice is nil (null)
	fmt.Println(nums)        // [] -> nill slice
	fmt.Println(nums == nil) // true
	fmt.Println(len(nums))   // 0

	// initialise while declaration:
	var nums2 = make([]int, 2) // type and initial size with 0 (for int) values. generally we have the initial value of the length is taken 0 to avoid the values with 0 without any use.
	fmt.Println(nums2)         // not nill but filled with zeroes
	fmt.Println(len(nums2))    // 2
	// capacity : maximum numbers of elements fit. (dynamic)
	fmt.Println(cap(nums2)) // capacity = 2 = initial length

	var nums3 = make([]int, 2, 5) // type and initial size and capacity.
	fmt.Println(nums3)

	// adding elements:
	nums3 = append(nums3, 1) // slice , element, addes in the end
	nums3 = append(nums3, 1) // slice , element, addes in the end
	nums3 = append(nums3, 1) // slice , element, addes in the end
	nums3 = append(nums3, 1) // slice , element, addes in the end
	nums3 = append(nums3, 1) // slice , element, addes in the end
	fmt.Println(nums3)
	fmt.Println(cap(nums3)) // capacity will be resized when the elements incereases ie. double to the current capacity. ie. 10
	// this capacity will remain 10 until no. of elements cross it.

	// another way to create slice:
	nums4 := []int{} // empty slice.
	fmt.Println(len(nums4))
	fmt.Println(cap(nums4))
	fmt.Println(nums4)

	// access using the index:
	// first make sure to give the length upto in the initialisation for the slice to have the value, to a count we want the index to be accessed.
	// nums4[0] = 9; // out of range
	fmt.Println(len(nums4))
	fmt.Println(cap(nums4))
	fmt.Println(nums4)

	// copy the slice  14:24
	var nums5 = make([]int, len(nums3)) // taking the nums5 same as the length of the nums3. so that the nums3 can be added/copied in the nums5.
	copy(nums5, nums3)                  // destination , source.
	// note : the slice in which copied is done should not be empty initial length. it's length should be some number to which you want the niumber to be copied.
	fmt.Println(nums3, nums5)

	// slice operator:
	var nums6 = []int{1, 2, 3}
	fmt.Println((nums6[0:2])) // slice ":" till the index 0 to 2 excluded the index last (2 here) value.
	//OR way:
	fmt.Println((nums6[:1])) // when not initial value given then it is from the start.
	fmt.Println((nums6[0:])) // similarly for the leaving the last value we will get till the end.

	// "slices" package:
	// various method inside, one of them is to check that two slices are equal or not.
	fmt.Println(slices.Equal(nums3, nums5))
	// this method compares one by one, one from each, then next.

	// 2-D slices:
	var nums7 = [][]int{{1, 2}, {3, 4, 5}}
	fmt.Println(nums7)

}
