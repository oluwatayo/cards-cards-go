package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	//card := newCard()
	//fmt.Println(card)
	cards := createSlice()
	cards = append(cards, "Six of Spade")

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
