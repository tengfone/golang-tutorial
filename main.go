package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	// var card string = "Ace of Spades"
	// Alternative way to declare a variable (Auto Infer)
	// card := "Ace of Spades"
	// card := newCard()

	// Declare a slice which requires data type and place object inside curly braces
	// card := []string{newCard(), newCard(), "3rd Object"}
	// card = append(card, "Six of Spades")
	// fmt.Println(card)

	// Using a new type of `type deck []string` from deck.go
	// newTypeCard := deck{"Ace of Diamonds", newCard()}
	// newTypeCard.print()

	cards := newDeck()
	hand, remainingDeck := deal(cards, 4)
	hand.print()
	remainingDeck.print()
}

// Must declare return type thus 'string'
func newCard() string {
	return "Five of Diamonds"
}

// Array in Go is a fixed length list of things while Slice is a list that can grow or shrink. [same data type]
