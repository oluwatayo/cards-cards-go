package main

func main() {
	//cards := newDeck()
	//hands, remainingCards := deal(cards, 5)
	// fmt.Println("Cards at hand after deal")
	// hands.print()
	// fmt.Println("Remaining cards after deal")
	// remainingCards.print()

	cards := newDeckFromFile("my_cards.txt")
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards.txt")
	cards.print()
}
