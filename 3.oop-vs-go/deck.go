package main

import (
	"fmt"
	"io/ioutil" // https://pkg.go.dev/io/ioutil -> ioutil is a subpackage inside the io package
	"strings"
)

type deck []string

// similuates a constructor, although Go is not OOP
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"spades", "diamonds", "hearts", "clubs"}
	cardsValues := []string{"ace", "two", "three", "four"}

	for _, suit := range cardSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// (d deck) is a receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// d is, by convention, a single letter identifier
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] // returns two values of type deck
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// https://pkg.go.dev/io/ioutil
// byte slice is a string
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
