package main

import "fmt"

func main() {
	cards := NewDeck()

	fmt.Println(cards.ToString())

	cards.WriteFile("my_file")
	oldCards := ReadDeckFromFile("my_file")

	fmt.Println(oldCards.ToString())

	fmt.Println(cards.ToString() == oldCards.ToString())
}
