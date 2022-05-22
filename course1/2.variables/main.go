package main

import "fmt"

func main() {
	var card string = "ace of spades"
	fmt.Println(card)

	// or

	var card2 = "ace of spades" // inferred
	fmt.Println(card2)

	// or

	card3 := "ace of spades" // := only for new assignments, not reassignments
	fmt.Println(card3)

	card4 := newCard()
	fmt.Println(card4)

	cards := []string{"ace of diamonds", newCard()}
	fmt.Println(cards)

	cards = append(cards, "six of spades")
	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "five of diamonds"
}
