package main

import "fmt"

func main() {
	tab := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, nbr := range tab {
		if nbr%2 == 0 {
			fmt.Print(nbr, ": Even\n")
		} else {
			fmt.Print(nbr, ": Odd\n")
		}
	}
}
