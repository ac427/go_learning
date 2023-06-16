package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}
type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	// printGreeting1(eb)
	// printGreeting2(sb)

}

func (eb englishBot) getGreeting() string {
	// VERY custom logic for generating english greeting
	return "Hello There"
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting1(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting2(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
