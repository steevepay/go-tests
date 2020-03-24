package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) == 0 {
		t.Errorf("Can't create a new deck: lenght different.")
	}

	if cards[len(cards)-1] != "K of Clubs" {
		t.Errorf("Can't create a new deck: last card different.")
	}
}

func TestShuffleDeck(t *testing.T) {
	initCards := newDeck()
	randCards := newDeck()
	randCards.shuffle()
	totalCards := len(randCards) - 1
	if initCards[0] == randCards[0] ||
		initCards[totalCards] == randCards[totalCards] {
		t.Errorf("Deck is not random")
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	fileName := "_decktesting"
	os.Remove(fileName)
	d := newDeck()
	d.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	if len(loadedDeck) != len(d) {
		t.Errorf("Error during savefile and loadfile.")
	}
}
