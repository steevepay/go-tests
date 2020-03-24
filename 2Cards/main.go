package main

func main() {

	// cards := newDeck()
	// cards.print()
	// hand, _ := deal(cards, 13)
	// hand.saveToFile("save.csv")
	cards := newDeckFromFile("save.csv")
	cards.shuffle()
	cards.print()
}
