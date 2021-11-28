package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
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
			cards = append(cards, value+" of "+suit)
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

func (d deck) WriteFile(filename string) {
	err := ioutil.WriteFile(filename, []byte(d.ToString()), 0666)

	if err != nil {
		panic(err)
	}
}

func ReadDeckFromFile(filename string) deck {
	fileByteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fileString := string(fileByteSlice)
	fileStringSlice := strings.Split(fileString, ",")
	return deck(fileStringSlice)
}

func (d deck) Suffle() {
	rand.Seed(time.Now().UnixNano())

	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
