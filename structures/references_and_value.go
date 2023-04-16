package main

import "fmt"

func slicing() {
	mySlice := []string{"a", "b", "c", "d", "e"}
	updateSlice(mySlice)

	// It will print ["XXX", "b", "c", "d", "e"], but why no need use pointer?
	fmt.Println(mySlice)

	// *The main reason why this is happening is because by default, a slice [] is a subset of an array, golang will
	// take the slice (which is actually a pointer to the array) and pass it to the function. Thus the name value vs reference.
	// Value Types: int, float, string, bool, structs
	// Reference Types: slices, maps, channels, pointers, functions
}

func updateSlice(s []string) {
	s[0] = "XXX"
}
