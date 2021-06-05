package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	ss := strings.Split(string(bs), ",")

	return deck(ss)
}
