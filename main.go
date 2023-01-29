package main


func main() {
	deck := newDeck()
  deck.shuffle()
  hand, deck := deal(deck, 5)
  hand.print()
}
