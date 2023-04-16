package main

import "fmt"

// Custom type called "bot", the function inside this interface is called "getGreeting" and it returns a string it is
// a 'promoted' function that uses the type bot. You can now use this interface as a type in your code.
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
	httpPackage()
}

func (englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// Very custom logic for generating a spanish greeting
	return "Hola!"
}

// Correct solution [ADD INTERFACE]
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Naive approach
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }
// Will fail cause cannot have 2 functions with same name
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
