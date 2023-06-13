package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
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

func (d deck) saveToFile() {
	message := []byte(strings.Join(d, "\n"))
	err := os.WriteFile("deck.txt", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (d deck) appendToFile() {
	f, err := os.OpenFile("deck_append.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(strings.Join(d, "\n")); err != nil {
		log.Println(err)
	}
	if _, err := f.WriteString("\n#########\n"); err != nil {
		log.Println(err)
	}
}

func (d deck) readFromFile() {
	data, err := os.ReadFile("deck_append.txt")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)
}
