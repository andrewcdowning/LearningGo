package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
// which is a slice of string
type deck []string

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Hearts", "Spades", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) saveToFile(fileName string) error {
	byteSlice := []byte(d.toString())
	return ioutil.WriteFile(fileName, byteSlice, 0666)
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) shuffle() {
  source := rand.NewSource(time.Now().UnixMilli())
  r := rand.New(source)

  for i := range d {
    randomIndex := r.Intn(len(d)-1)
    d[i], d[randomIndex] = d[randomIndex], d[i]
    
  }
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}


