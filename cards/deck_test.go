package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 12 {
		t.Errorf("Expected deck length of 12, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Three of Clubs" {
		t.Errorf("Expected last card of Three of Clubs, but got %v", d[len(d)-1])
	}
}

// Test file save and load
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 12 {
		t.Errorf("Expected 12 cards in deck, got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
