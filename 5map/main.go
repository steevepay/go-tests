package main

import "fmt"

type colors map[string]string

func (c colors) printMap() {
	for key, value := range c {
		fmt.Printf("%v %v\n", key, value)
	}
}

func main() {
	c := colors{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"green": "#00ff00",
		"white": "#ffffff",
	}
	// Working too:
	// var colors = make(map[string]string)
	// colors["white"] = "#ffffff"
	// Delete an element
	// delete(colors, "white")
	c.printMap()
	// fmt.Print(c)
}
