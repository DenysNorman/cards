package main

func main() {
	// create a deck and save to file

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	//export a deck from file

	// cards := newDeckFromFile("my_cardss")
	// cards.print()

	//shuffle a deck

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
