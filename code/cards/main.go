package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 3)

	hand.Print()
	remainingCards.Print()
}
