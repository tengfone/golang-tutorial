// Create test file for golang is default xxx_test.go and run with `go test` command
package main

import (
	"os"
	"testing"
)

// By default you need to put testing.T for test. t is test handler
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 12 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Three of Clubs" {
		t.Errorf("Expected last card of Three of Clubs, but got %v", d[len(d)-1])
	}
}

// Long name is ok cause its meant for human readable
func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 12 {
		t.Errorf("Expected 12 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
