package main

import "fmt"

type bot interface {
	getGreating() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreating(eb)
	printGreating(sb)
}

func printGreating(b bot) {
	fmt.Println(b.getGreating())
}

func (englishBot) getGreating() string {
	// different logic
	return "Hi there!"
}

func (eb spanishBot) getGreating() string {
	// different logic
	return "Hola!"
}
