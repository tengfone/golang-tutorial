package main

import "fmt"

func main() {
	integers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, integer := range integers {
		if integer%2 == 0 {
			fmt.Printf("%v is even\n", integer)
		} else {
			fmt.Printf("%v is odd\n", integer)
		}
	}
}
