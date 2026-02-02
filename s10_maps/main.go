package main

import (
	"fmt"
	"maps"
)

// maps: associate data structure eg. hash in java.

func main() {
	m := make(map[string]string) // map[keytype]valuetype
	var m2 map[int]int = make(map[int]int)

	// setting value in map
	m["apple"] = "rahul"
	m["backend"] = "golang"
	m2[2] = 2

	// get a value on the key
	fmt.Println(m["apple"])
	fmt.Println(m["backend"])
	fmt.Println(m["phone"]) // if the key value does not exitst then it returns the empty value of the datatype of the value. eg. "" for string, 0 for int, etc.

	//map length: no. of keys in the map.
	fmt.Println(len(m))

	// delete a key from the map.
	delete(m, "backend") // map, key
	fmt.Println(m)       // checking

	// clear : to completely empty map.
	clear(m)
	fmt.Println(m)

	// other way to create the map: used when we have the elements already.
	m3 := map[string]int{"apple": 31, "spain": 12} // similar to JS object
	fmt.Println(m3)

	// to check a key exist in map or not.
	value, boolean := m3["apples"]
	// boolean tells the value exist in the map or not. we can use it with if.
	// value if key exist return the value for the key.
	if boolean {
		fmt.Println(value)
	} else {
		fmt.Println("value does not exist")
	}

	// check two maps are equal or not: using the "maps" package
	m4 := map[string]int{"apple": 31, "spain": 12}
	fmt.Println(maps.Equal(m3, m4)) 
	// note: only same keytype and valuetype maps can be compared.

}
