package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	hand, cards := deal(cards, 2)
	fmt.Println("Your hand contains: ")
	hand.print()
	fmt.Println("Cards left in the deck are: ")
	cards.print()
}
