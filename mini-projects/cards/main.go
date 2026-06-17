package main

import "fmt"

func main() {
	cards := newDeckFromFile("my_cards.txt")
	fmt.Println("The length of the deck is: ", len(cards))

	cards.shuffle()

	hand, remainingCards := deal(&cards, 5)

	fmt.Println("In Hand: ")
	hand.print()

	fmt.Println("Remaining Cards: ")
	remainingCards.print()

	// Run the following line to save the cards to a file
	// cards.saveToFile("my_cards.txt")
}
