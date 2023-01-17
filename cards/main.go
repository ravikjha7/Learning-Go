package main

// import "fmt"

func main() {
	
	// var card string = "Ace of Spades"

	// card := newCard()
	// card = "Five of Diamonds"

	// fmt.Println(card)

	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards := newDeck()

	// hand, _ := deal(cards, 5)

	// // hand.print()
	// // remainingCards.print()

	// fmt.Println(hand.toString())

	// hand.saveToFile("NewHand")

	cards := newDeckFromFile("NewHand")
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}