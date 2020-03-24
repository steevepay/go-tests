package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of "deck"
// Which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Ace", "Diamonds", "Hearts", "Clubs"}
	cardValue := []string{"One", "Two", "Tree", "Four", "Five", "Six", "Seven", "Height", "Nine", "J", "Q", "K"}

	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Receiver (d deck) any variable of type deck now have to access to print method.
// I can use d is a instance of the variable deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	// cards := ""
	// for _, card := range d {
	// 	cards += card + "\n"
	// }
	return strings.Join(d, ",")
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(fileName string) {
	buffer := []byte(d.toString())
	err := ioutil.WriteFile(fileName, buffer, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func newDeckFromFile(fileName string) deck {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return strings.Split(string(content), ",")
}
