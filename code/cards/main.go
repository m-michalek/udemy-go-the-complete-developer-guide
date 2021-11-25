package main

import "fmt"

func main() {
	cards := NewDeck()

	fmt.Println(cards.ToString())
}
