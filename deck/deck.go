package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Diamonds", "Clubs", "Spade", "Heart"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suite+" of "+value)
		}
	}
	return cards

}

func (d deck) print() {
	for i, value := range d {
		fmt.Println(i, value)
	}
}
