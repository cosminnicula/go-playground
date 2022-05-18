package main

import "fmt"

// go run main.go deck.go
func main() {
	cards := newDeck()

	// cards.print()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	// type conversion
	greeting := "hi there"
	fmt.Println([]byte(greeting))

	// toString
	fmt.Println(cards.toString())

	cards.saveToFile("cards.txt")
}
