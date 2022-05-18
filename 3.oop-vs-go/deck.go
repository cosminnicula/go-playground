package main

import "fmt"

type deck []string

// (d deck) is a receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
