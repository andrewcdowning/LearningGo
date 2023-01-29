package main

import (
	"os"
	"testing"
)


func TestNewDeck(t *testing.T) {
  d := newDeck()
  
  if len(d) != 52 {
    t.Error(len(d))
  }
}

func TestFirstCar(t *testing.T) {
  d := newDeck()
  if d[0] != "Ace of Hearts" {
    t.Errorf("Expected first card to be Ace of Hearts, but got %v", d[0])
  }
}

func TestDealHand(t *testing.T) {
  d := newDeck()
  hand, deck := deal(d, 5)
  if len(hand) != 5 {
    t.Errorf("Expected hand of 5 cards, but got %v", len(hand))
  }
  if len(deck) != 47 {
    t.Errorf("Expected deck to be 47 cards, but got %v", len(deck))
  }
}

func TestShuffleDeck(t *testing.T) {
  d:=newDeck()
  d.shuffle()
  d1:=newDeck()
  if d[0] == d1[0] {
    t.Errorf("Deck was not shuffled")
  }
}

func TestSaveDeckToFile(t *testing.T) {
  os.Remove("_deckTesting")
  d := newDeck()
  d.saveToFile("_deckTesting")
  loadedDeck := newDeckFromFile("_deckTesting")

  if len(loadedDeck) != 52 {
    t.Errorf("Unexpected deck length")

  }
  os.Remove("_deckTesting")
}

