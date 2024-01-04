package main

import "fmt"

func main() {
	//cards := newDeck()
	//hands, remainingCards := deal(cards, 5)
	// fmt.Println("Cards at hand after deal")
	// hands.print()
	// fmt.Println("Remaining cards after deal")
	// remainingCards.print()

	cards := newDeck()
	fmt.Println(cards.toString())
}
