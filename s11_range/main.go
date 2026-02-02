package main

import "fmt"

// range : used for iteration over the data-structure
func main() {
	nums := []int{6, 7, 8}

	// one way : for loop to iterate
	for i := 0; i < 3; i++ {
		fmt.Println(nums[i])
	}

	// 2nd way : range
	var sum int
	for index, num := range nums {
		sum += num
		fmt.Println(index, " : ", num)
	}
	fmt.Println(sum)

	// 3rd way : map iteration
	m := make(map[string]int)
	m["apple"] = 0
	m["mango"] = 3

	for k, v := range m {
		fmt.Println(k, v) // key , value
	}
	// though we can print only keys as well.

	// 4th way : range over the string. this prints the unicode of the each character ie. unicode of the rune.
	for i, c := range "golang" {
		fmt.Println("starting byte of rune", i, " character unicode : ", c, string(c))
	}
	// rune : self study
	// string() : method converts the unicode to character. 
}
