package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }
	// OR
	// Another way of declaring maps
	// var colors map[string]string
	// OR
	colors := make(map[string]string)

	// Add new key value pair
	colors["white"] = "#ffffff"
	colors["black"] = "#000000"
	colors["red"] = "#ff0000"

	// Delete a key value pair
	delete(colors, "white")

	// Iterate through maps
	iterateMap(colors)

	fmt.Println(colors)
}

func iterateMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}
