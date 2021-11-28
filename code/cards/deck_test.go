package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	underTest := NewDeck()

	if len(underTest) != 16 {
		t.Errorf("Expected deck of length 16 but got %v", len(underTest))
	}

	if underTest[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", underTest[0])
	}

	if underTest[len(underTest)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs but got %v", underTest[len(underTest)-1])
	}
}

func TestWriteFileAndReadDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := NewDeck()

	deck.WriteFile("_decktesting")

	underTest := ReadDeckFromFile("_decktesting")

	if len(underTest) != 16 {
		t.Errorf("Expected deck of length 16 but got %v", len(underTest))
	}

	os.Remove("_decktesting")
}
