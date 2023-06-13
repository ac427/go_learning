package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func (d deck) shuffle1() {

	for i := range d {
		randomNumber := rand.Intn(len(d) - 1)
		d[i], d[randomNumber] = d[randomNumber], d[i]
	}

}

func (d deck) shuffle2() {
	for i := range d {
		seedVal := time.Now().UnixNano()
		rand.Seed(seedVal)
		randomNumber := rand.Intn(len(d) - 1)
		d[i], d[randomNumber] = d[randomNumber], d[i]
	}
}

func (d deck) shuffle3() {
	for i := range d {
		seedVal := time.Now().UnixNano()
		source := rand.NewSource(seedVal)
		r := rand.New(source)
		randomNumber := r.Intn(len(d) - 1)
		d[i], d[randomNumber] = d[randomNumber], d[i]
	}
}
