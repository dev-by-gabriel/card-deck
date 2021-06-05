package main

import "fmt"

// Create a new type of deck
// Which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardValues := []string{"A", "2", "3", "4", "5", "6",
		"7", "8", "9", "10", "J", "Q", "K"}
	cardSuits := []string{"S", "D", "H", "C"}

	for _, value := range cardValues {
		for _, suit := range cardSuits {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
