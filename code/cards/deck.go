package main

import (
	"fmt"
	"strings"
)

type deck []string

func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func NewDeck() deck {
	var cards deck

	cardSuits := []string{
		"Spades",
		"Diamonds",
		"Hearts",
		"Clubs",
	}
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func Deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) ToString() string {
	return strings.Join(d, ",")
}
