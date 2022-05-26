package main

import "fmt"

// go run main.go deck.go
func main() {
	cards := newDeck()

	// cards.print()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	fmt.Println("***")

	// type conversion
	greeting := "hi there"
	fmt.Println([]byte(greeting))

	fmt.Println("***")

	// toString
	fmt.Println(cards.toString())

	cards.saveToFile("cards.txt")

	fmt.Println("***")

	cardsFromFile := newDeckFromFile("cards.txt")
	fmt.Println(cardsFromFile.toString())

	fmt.Println("***")

	cardsFromFile.shuffle()
	fmt.Println(cardsFromFile.toString())
}
