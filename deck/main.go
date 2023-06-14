package main

//import "fmt"

func main() {
	d := newDeck()
	// fmt.Println(d)
	// d.print()
	d.shuffle1()
	hand, remaining := deal(d, 3)
	hand.print()
	remaining.print()
	// fmt.Println("###############")
	// d.print()
	//fmt.Println("###############")
	d.shuffle2()
	//d.print()
	//fmt.Println("###############")
	d.shuffle3()
	//d.print()
	d.saveToFile("deck.txt")
	d.appendToFile()
	//d.readFromFile()
}
