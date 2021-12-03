package main

func main() {
	cards := NewDeck()

	cards.WriteFile("my_file")
	cardsFomFile := ReadDeckFromFile("my_file")

	cardsFomFile.Suffle()
	cardsFomFile.Print()
}
